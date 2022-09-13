package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"snapp-featureflag/internal/package/config"
)

type Service interface {
	SetByKey(ctx context.Context, key string, value string) error
	GetByKey(ctx context.Context, key string) (string, error)
	DeleteByKey(ctx context.Context, key string) error
	GetAllByKey(ctx context.Context, keyPattern string) ([]string, error)
	GetList(ctx context.Context, key string) ([]string, error)
	AddToList(ctx context.Context, key string, values ...string) error
}

type ServiceImpl struct {
	redisClient *redis.Client
}

func NewCacheService(config *config.AppConfig) (ServiceImpl, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(config.Redis.Host, ":", config.Redis.Port),
		Password: config.Redis.Password,
	})
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return ServiceImpl{}, err
	}

	return ServiceImpl{redisClient}, nil
}

func (c ServiceImpl) SetByKey(ctx context.Context, key string, value string) error {
	return c.redisClient.Set(ctx, key, value, 0).Err()
}

func (c ServiceImpl) GetByKey(ctx context.Context, key string) (string, error) {
	return c.redisClient.Get(ctx, key).Result()
}

func (c ServiceImpl) DeleteByKey(ctx context.Context, key string) error {
	return c.redisClient.Del(ctx, key).Err()
}

func (c ServiceImpl) GetAllByKey(ctx context.Context, keyPattern string) ([]string, error) {
	return c.redisClient.Keys(ctx, keyPattern).Result()
}

func (c ServiceImpl) GetList(ctx context.Context, key string) ([]string, error) {
	return c.redisClient.LRange(ctx, key, 0, -1).Result()
}

func (c ServiceImpl) AddToList(ctx context.Context, key string, values ...string) error {
	return c.redisClient.LPush(ctx, key, values).Err()
}
