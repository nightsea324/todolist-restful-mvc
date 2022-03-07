package redis

import (
	"context"
	"time"
)

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
