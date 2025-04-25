package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis")
}

func SetCache(key string, value string, expiration time.Duration) {
	RedisClient.Set(ctx, key, value, expiration)
}

func GetCache(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func DeleteCache(key string) {
	RedisClient.Del(ctx, key)
}

func UpdateCache(key string, value string, expiration time.Duration) {
	RedisClient.Del(ctx, key)
	RedisClient.Set(ctx, key, value, expiration)
}
