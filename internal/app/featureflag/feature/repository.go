package feature

import (
	"context"
	"encoding/json"
	"fmt"
	redis2 "github.com/go-redis/redis/v9"
	"snapp-featureflag/internal/app/featureflag/feature/entity"
	"snapp-featureflag/internal/package/service/cache"
)

type Repository interface {
	GetFeature(ctx context.Context, featureName string) (*entity.Feature, error)
	CreateFeature(ctx context.Context, feature *entity.Feature) error
	UpdateFeature(ctx context.Context, featureName string, feature *entity.Feature) error
	DeleteFeature(ctx context.Context, featureName string) error
	GetFeatureFlagsByUserIfExist(ctx context.Context, userId uint) ([]*entity.FeatureWithFlag, error)
	SetFeatureFlagsByUser(ctx context.Context, userId uint, features []*entity.FeatureWithFlag) error
	GetAllFeatures(ctx context.Context) ([]*entity.Feature, error)
}

type RepositoryImpl struct {
	cache cache.Service
}

func NewRepository(cache cache.Service) RepositoryImpl {
	return RepositoryImpl{cache}
}

func (r RepositoryImpl) GetFeature(ctx context.Context, featureName string) (*entity.Feature, error) {
	featureStr, err := r.cache.GetByKey(ctx, r.getFeatureCacheKey(featureName))
	if err != nil {
		if err == redis2.Nil {
			return nil, nil
		}
		return nil, err
	}

	feature := &entity.Feature{}
	if err = json.Unmarshal([]byte(featureStr), &feature); err != nil {
		return nil, err
	}

	return feature, nil

}

func (r RepositoryImpl) CreateFeature(ctx context.Context, feature *entity.Feature) error {
	ft, err := json.Marshal(feature)
	if err != nil {
		return err
	}
	if err = r.cache.SetByKey(ctx, r.getFeatureCacheKey(feature.Name), string(ft)); err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) UpdateFeature(ctx context.Context, featureName string, feature *entity.Feature) error {
	ft, err := json.Marshal(feature)
	if err != nil {
		return err
	}
	if err = r.cache.SetByKey(ctx, r.getFeatureCacheKey(featureName), string(ft)); err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) DeleteFeature(ctx context.Context, featureName string) error {
	if err := r.cache.DeleteByKey(ctx, r.getFeatureCacheKey(featureName)); err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) SetFeatureFlagsByUser(ctx context.Context,
	userId uint, features []*entity.FeatureWithFlag) error {
	featuresJsons := make([]string, len(features), len(features))

	for i, feature := range features {
		ftJson, err := json.Marshal(feature)
		if err != nil {
			return err
		}
		featuresJsons[i] = string(ftJson)
	}

	if err := r.cache.DeleteByKey(ctx, r.getUserCacheKey(userId)); err != nil {
		return err
	}

	if len(featuresJsons) != 0 {
		if err := r.cache.AddToList(ctx, r.getUserCacheKey(userId), featuresJsons...); err != nil {
			return err
		}
	}

	return nil
}

func (r RepositoryImpl) GetFeatureFlagsByUserIfExist(
	ctx context.Context, userId uint) ([]*entity.FeatureWithFlag, error) {
	featuresJsonList, err := r.cache.GetList(ctx, r.getUserCacheKey(userId))
	features := make([]*entity.FeatureWithFlag, len(featuresJsonList), len(featuresJsonList))

	if err != nil {
		if err == redis2.Nil {
			return features, nil
		}
		return nil, err
	}

	for i, featureJson := range featuresJsonList {
		feature := &entity.FeatureWithFlag{}
		if err = json.Unmarshal([]byte(featureJson), &feature); err != nil {
			return nil, err
		}
		features[i] = feature
	}

	return features, nil

}

func (r RepositoryImpl) GetAllFeatures(ctx context.Context) ([]*entity.Feature, error) {
	featuresNames, err := r.cache.GetAllByKey(ctx, "feature::*")
	features := make([]*entity.Feature, len(featuresNames), len(featuresNames))
	if err != nil {
		return features, err
	}

	for i, name := range featuresNames {
		feature := &entity.Feature{}
		featureStr, err := r.cache.GetByKey(ctx, name)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal([]byte(featureStr), &feature); err != nil {
			return nil, err
		}
		features[i] = feature
	}

	return features, nil
}

func (r RepositoryImpl) getFeatureCacheKey(featureName string) string {
	return fmt.Sprintf("feature::%s", featureName)
}

func (r RepositoryImpl) getUserCacheKey(userId uint) string {
	return fmt.Sprintf("user::%d", userId)
}
