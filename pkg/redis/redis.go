package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func GetRedisConnection() *redis.Client {
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		redisHost = "127.0.0.1"
	}

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		redisPort = "6379"
	}

	redisAddr := fmt.Sprintf("%s:%s", redisPort, redisHost)
	conn := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return conn
}
