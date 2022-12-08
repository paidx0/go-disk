package models

import (
	"code/define"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()
var Rdb = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine(define.DBEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", define.DBUser, define.DBPasswd, define.DBHost, define.DBPort, define.DBName))
	if err != nil {
		log.Fatalf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     define.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
