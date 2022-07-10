package cache

import (
	"github.com/go-redis/redis/v8"
)

func NewClient(addr, pw string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       0,
	})
}
