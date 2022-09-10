package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"snapp-featureflag/internal/package/config"
)

type CacheService interface {
	SetByKey(ctx context.Context, key string, value string) error
	GetByKey(ctx context.Context, key string) (string, error)
	GetAllByKey(ctx context.Context, keyPattern string) ([]string, error)
	GetList(ctx context.Context, key string) ([]string, error)
	AddToList(ctx context.Context, key string, values ...string) error
}

type CacheServiceImpl struct {
	redisClient *redis.Client
}

func NewCacheService(config *config.AppConfig) (CacheServiceImpl, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(config.Redis.Host, ":", config.Redis.Port),
		Password: config.Redis.Password,
	})
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return CacheServiceImpl{}, err
	}

	return CacheServiceImpl{redisClient}, nil
}

func (c CacheServiceImpl) SetByKey(ctx context.Context, key string, value string) error {
	return c.redisClient.Set(ctx, key, value, 0).Err()
}

func (c CacheServiceImpl) GetByKey(ctx context.Context, key string) (string, error) {
	return c.redisClient.Get(ctx, key).Result()
}

func (c CacheServiceImpl) GetAllByKey(ctx context.Context, keyPattern string) ([]string, error) {
	return c.redisClient.Keys(ctx, keyPattern).Result()
}

func (c CacheServiceImpl) GetList(ctx context.Context, key string) ([]string, error) {
	return c.redisClient.LRange(ctx, key, 0, -1).Result()
}

func (c CacheServiceImpl) AddToList(ctx context.Context, key string, values ...string) error {
	return c.redisClient.LPush(ctx, key, values).Err()
}
