package feature

import (
	"context"
	"errors"
	"fmt"
	"snapp-featureflag/internal/app/featureflag/feature/entity"
)

type QueryHandler struct {
	Repository
}

func NewQueryHandler(repository Repository) QueryHandler {
	return QueryHandler{repository}
}

func (q QueryHandler) Get(ctx context.Context, query GetFeatureQuery) (*GetFeatureQueryResponse, error) {
	if query.Name == "" {
		return nil, errors.New("feature name can't be nil")
	}

	feature, err := q.Repository.GetFeature(ctx, query.Name)
	if err != nil {
		return nil, err
	}

	if feature == nil {
		return nil, errors.New(
			fmt.Sprintf("feature with name \" %s \" does not exist", query.Name))
	}

	resp := &GetFeatureQueryResponse{
		Name:     feature.Name,
		Coverage: feature.Rule.Coverage,
	}

	if feature.Rule.MinVersion != nil {
		resp.MinVersion = feature.Rule.MinVersion.ToString()
	}
	return resp, nil

}

func (q QueryHandler) GetActiveFeaturesByUserId(ctx context.Context,
	query GetActiveFeaturesByUserIdQuery) (*GetActiveFeaturesByUserIdQueryResponse, error) {
	version, err := entity.NewValidVersionFromString(query.Version)
	if err != nil {
		return nil, err
	}
	user, err := entity.NewUser(query.UserId, *version)
	if err != nil {
		return nil, err
	}

	previousActiveFeatures, err := q.Repository.GetFeatureFlagsByUserIfExist(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	features, err := q.Repository.GetAllFeatures(ctx)
	if err != nil {
		return nil, err
	}

	activeFeatures := user.GetFeatureFlags(previousActiveFeatures, features)

	if err = q.Repository.SetFeatureFlagsByUser(ctx, user.Id, activeFeatures); err != nil {
		return nil, err
	}

	return mapFeaturesToActiveFeaturesResp(activeFeatures), nil

}

func mapFeaturesToActiveFeaturesResp(
	features []*entity.FeatureWithFlag) *GetActiveFeaturesByUserIdQueryResponse {
	resp := &GetActiveFeaturesByUserIdQueryResponse{}
	names := make([]FeatureName, 0)

	for _, feature := range features {
		if feature.IsActive {
			names = append(names, FeatureName{feature.Name})
		}
	}

	resp.FeaturesNames = names

	return resp
}
