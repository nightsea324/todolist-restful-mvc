package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

// NewClient - 初始化redis
func NewClient() {

	var ctx = context.TODO()

	// 連線redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 確認連線
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
