package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // เช่น "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // ไม่มี password ให้ใช้ ""
		DB:       0,                           // ใช้ default DB
	})
}

func GetRedisClient() *redis.Client {
	return client
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}
