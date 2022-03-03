package redis

import (
	"context"
	"fmt"
	"time"

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

// Set - 寫入redis
func Set(ctx context.Context, key string, value interface{}) error {

	if err := client.Set(ctx, key, value, time.Hour*24*30).Err(); err != nil {
		return err
	}
	return nil
}

// Get - 讀取redis
func Get(ctx context.Context, key string) (string, error) {

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
