package feature

import (
	"context"
	"errors"
	"snapp-featureflag/internal/app/featureflag/feature/entity"
)

type CommandHandler struct {
	Repository
}

func NewCommandHandler(repository Repository) CommandHandler {
	return CommandHandler{repository}
}

func (c CommandHandler) Create(ctx context.Context,
	command CreateFeatureCommand) error {

	newFeature, err := entity.NewFeatureFromParams(entity.CreateFeatureParams{
		Name:       command.Name,
		Coverage:   command.Coverage,
		MinVersion: command.MinVersion,
	})
	if err != nil {
		return err
	}

	feature, err := c.Repository.GetFeature(ctx, newFeature.Name)
	if err != nil {
		return err
	}

	if feature != nil {
		return errors.New("feature already exists")
	}

	if err = c.Repository.CreateFeature(ctx, newFeature); err != nil {
		return err
	}

	return nil

}

func (c CommandHandler) Update(ctx context.Context,
	command UpdateFeatureCommand) error {

	newFeature, err := entity.NewFeatureFromParams(entity.CreateFeatureParams{
		Name:       command.Name,
		Coverage:   command.Coverage,
		MinVersion: command.MinVersion,
	})
	if err != nil {
		return err
	}

	feature, err := c.Repository.GetFeature(ctx, newFeature.Name)
	if err != nil {
		return err
	}

	if feature == nil {
		return errors.New("feature does not exist")
	}

	if err := c.updateFeature(ctx, feature, newFeature); err != nil {
		return err
	}

	return nil

}

func (c CommandHandler) updateFeature(ctx context.Context,
	feature *entity.Feature, newFeature *entity.Feature) error {
	if err := feature.Update(newFeature); err != nil {
		return err
	}
	if err := c.Repository.UpdateFeature(ctx, feature.Name, newFeature); err != nil {
		return err
	}
	return nil
}

func (c CommandHandler) Delete(ctx context.Context, command DeleteFeatureCommand) error {
	if command.Name == "" {
		return errors.New("feature name can't be nil")
	}

	if err := c.Repository.DeleteFeature(ctx, command.Name); err != nil {
		return err
	}
	return nil

}
