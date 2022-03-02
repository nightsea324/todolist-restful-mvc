package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func NewClient() {
	var ctx = context.TODO()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}

func Set(ctx context.Context, key string, value interface{}) error {
	if err := client.Set(ctx, key, value, time.Hour*24*30).Err(); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, key string) string {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("client Get failed", err)
	}
	return value
}
