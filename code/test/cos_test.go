package test

import (
	"code/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUpload_localfile(t *testing.T) {
	u, _ := url.Parse(define.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	key := "uploads/test01.jpg"

	_, _, err := client.Object.Upload(context.Background(), key, "./00.png", nil)
	if err != nil {
		panic(err)
	}
}

func TestFileUpload_putfile(t *testing.T) {
	u, _ := url.Parse(define.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	key := "uploads/test2.jpg"

	f, err := os.Open("./00.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(context.Background(), key, f, nil)
	if err != nil {
		panic(err)
	}
}
