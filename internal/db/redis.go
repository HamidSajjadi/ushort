package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis(address, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	err := ping(rdb)
	return rdb, err

}

func ping(client *redis.Client) error {
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
