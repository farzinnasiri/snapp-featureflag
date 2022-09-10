package feature

import (
	"context"
	"snapp-featureflag/internal/package/service/redis"
)

type Repository interface {
	GetFeature(ctx context.Context, featureName string) error
	CreateFeature(ctx context.Context, feature Feature) error
	UpdateFeature(ctx context.Context, featureName string, feature Feature) error
	DeleteFeature(ctx context.Context, featureName string) error
}

type RepositoryImpl struct {
	cache redis.CacheService
}

func NewRepository(cache redis.CacheService) RepositoryImpl {
	return RepositoryImpl{cache}
}

func (RepositoryImpl) GetFeature(ctx context.Context, featureName string) error {
	panic("implement me")
}

func (RepositoryImpl) CreateFeature(ctx context.Context, feature Feature) error {
	panic("implement me")
}

func (RepositoryImpl) UpdateFeature(ctx context.Context, featureName string, feature Feature) error {
	panic("implement me")
}

func (RepositoryImpl) DeleteFeature(ctx context.Context, featureName string) error {
	panic("implement me")
}
