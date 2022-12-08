package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "go-disk-paidx0"

var TokenExpire int64 = 8 * 60 * 60 // token期限 8个小时

var MailPassword = "xxxxxxxx" // 邮箱授权码，写自己的
var CodeExpire = time.Second * 200    // 验证码失效时间 200秒

var DBEngine = "mysql"
var DBUser = "root"
var DBPasswd = "root"
var DBPort = "3306"
var DBHost = "127.0.0.1"
var DBName = "cloud-disk"

var RedisAddr = "localhost:6379"

var BucketName = "https://xxxxxxxxxxxxxxxxxxxxxxxxx.myqcloud.com" // 腾讯云存储仓库，写自己的
var SecretId = "xxxxxxxxxxxxxxxx"
var SecretKey = "xxxxxxxxxxxxxxx"

var PageSize = 15 // 默认查询一页的容量

var Datetime = "2006-01-02 15:04:05" // 时间格式
