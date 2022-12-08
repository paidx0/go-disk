package utils

import (
	"bytes"
	"code/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func CreateToken(id int, identity, name string, expire int64) (string, error) {
	uc := define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * time.Duration(expire)).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := claims.Claims.(*define.UserClaim); ok && claims.Valid {
		return uc, nil
	} else {
		return nil, errors.New("token失效的！！！")
	}
}

func MailCodeSend(mail, code string) error {
	e := email.NewEmail()
	e.From = "Go-disk <paidxing0@163.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "paidxing0@163.com", define.MailPassword, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func MailCodeCreate() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func UUID() string {
	return uuid.NewV4().String()
}

func FileUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "uploads/" + UUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return "", err
	}
	return key, nil
}

func FileChunkUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "uploads/" + UUID() + path.Ext(fileHeader.Filename)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	// 分块初始化拿到id
	name := key
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		return "", err
	}
	UploadID := v.UploadID

	// 上传分块，块数最多10000块
	chunkObject := make([]cos.Object, 0)
	chunkLen := fileHeader.Size / 100 // 100块，每一块的长度

	// todo 不分块就 1块的时候上传没问题，分块就无法响应，留个坑
	for i := 0; i < 100; i++ {
		partNum := i + 1
		resp, err := client.Object.UploadPart(
			context.Background(), key, UploadID, partNum, bytes.NewReader(buf.Bytes()[int64(i)*chunkLen:int64(i+1)*chunkLen]), nil,
			//context.Background(), key, UploadID, partNum, bytes.NewReader(buf.Bytes()), nil,
		)
		if err != nil {
			return "", err
		}
		partETag := strings.Trim(resp.Header.Get("ETag"), "\"") // 分块唯一标识，要去掉两头的引号
		chunkObject = append(chunkObject, cos.Object{
			PartNumber: partNum,
			ETag:       partETag,
		})
	}

	// 分块碎片合并
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, chunkObject...)
	_, _, err = client.Object.CompleteMultipartUpload(context.Background(), key, UploadID, opt)
	if err != nil {
		return "", err
	}
	return key, nil
}

func FileDownload(path, name, ext string) error {
	u, _ := url.Parse(define.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})
	key := path

	// 创建 "./download/" 文件夹
	os_path, _ := os.Getwd()
	pathDir := os_path + "/download/"
	if !PathExists(pathDir) {
		os.Mkdir("./download", os.ModePerm)
	}
	// 下载到本地文件
	_, err := client.Object.GetToFile(context.Background(), key, pathDir+name+ext, nil)
	if err != nil {
		return err
	}
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
