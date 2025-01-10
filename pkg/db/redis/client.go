package redis

import "github.com/redis/go-redis/v9"

func Init(addr string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // 没有密码，默认值
		DB:       db,       // 默认DB 0
	})

	return rdb
}
