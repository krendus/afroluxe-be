package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func connectRedis() *redis.Client {
	opt, err := redis.ParseURL(env.RedisUrl)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}

var r = connectRedis()

func SetRedisValue(key string, value interface{}) error {
	ctx := context.Background()
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expireAt := time.Hour

	return r.Set(ctx, key, p, expireAt).Err()
}

func GetRedisValue(key string, dest interface{}) error {
	ctx := context.Background()
	val, err := r.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func DelRedisValue(key string) error {
	ctx := context.Background()
	return r.Del(ctx, key).Err()
}
