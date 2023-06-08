package cache

import (
	"context"
	"envoy-golang-filter-hub/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Redis *redis.Client

func Init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Get().Redis.Address,
		Password: config.Get().Redis.Password,
		DB:       config.Get().Redis.DB,
		//超时
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  3 * time.Second,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Println("connect Redis failed")
		panic(err)
	}
	log.Println("connect Redis success")
}
