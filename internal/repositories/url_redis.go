package repositories

import (
	"context"
	"encoding/json"
	"github.com/HamidSajjadi/ushort/internal"
	"github.com/go-redis/redis/v8"
)

type RedisURLRepo struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisRepo(rdb *redis.Client) *RedisURLRepo {
	ctx := context.Background()
	return &RedisURLRepo{
		rdb: rdb,
		ctx: ctx,
	}
}

func (r RedisURLRepo) GetOne(shortenedURL string) (*URLModel, error) {
	val, err := r.rdb.Get(r.ctx, shortenedURL).Result()
	if err != nil {
		if err == redis.Nil {
			err = internal.NotFoundErr
		}
		return nil, err
	}
	var url URLModel
	err = json.Unmarshal([]byte(val), &url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

//Save if url does not exist already, creates a URLModel, json stringify it so that it can be saved into redis
func (r RedisURLRepo) Save(sourceURL string, shortURL string) (url *URLModel, err error) {

	url = &URLModel{
		Source:    sourceURL,
		Shortened: shortURL,
		Views:     0,
	}
	err = r.saveModel(url)
	return url, err
}

func (r RedisURLRepo) IncViews(shortenedURL string) error {
	urlModel, err := r.GetOne(shortenedURL)
	if err != nil {
		return err
	}
	urlModel.Views += 1
	return r.saveModel(urlModel)
}

func (r RedisURLRepo) saveModel(url *URLModel) error {
	val, err := json.Marshal(url)
	if err != nil {
		return err
	}
	err = r.rdb.Set(r.ctx, url.Shortened, val, 0).Err()
	return err
}
