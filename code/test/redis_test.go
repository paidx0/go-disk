package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "test", "value", time.Second*60).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}

func TestGetValue_2(t *testing.T) {
	val2, err := rdb.Get(ctx, "abc").Result()
	if err != nil {
		t.Error("空")
	}
	t.Log(val2)
}
