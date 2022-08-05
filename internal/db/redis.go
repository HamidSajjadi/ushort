package db

import "github.com/go-redis/redis/v8"

func NewRedis(address, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	return rdb

}
