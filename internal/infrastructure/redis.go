package infrastructure

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	RedisClient = rdb
	return rdb
}
