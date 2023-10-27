package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/DanielDDHM/world-coin-converter/pkg/redis"
)

var cache = redis.GetRedisConnection()

var ctx = context.Background()

func Set(key string, value string, ttl time.Duration) {
	cache.Set(ctx, key, value, ttl*time.Second)
}

func Get(key string) (interface{}, error) {
	data, err := cache.Get(ctx, key).Bytes()

	if err != nil {
		return nil, err
	}

	var jsonMap interface{}
	json.Unmarshal([]byte(string(data)), &jsonMap)
	return jsonMap, nil
}

func Del(key string) {
	cache.Del(ctx, key)
}
