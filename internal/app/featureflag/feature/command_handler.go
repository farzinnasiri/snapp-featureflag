package feature

import "context"

type CommandHandler struct {
	Repository
}

func NewCommandHandler(repository Repository) CommandHandler {
	return CommandHandler{repository}
}

func (c CommandHandler) CreateFeature(ctx context.Context,
	command CreateFeatureCommand) error {

	return nil

}

func (c CommandHandler) UpdateFeature(ctx context.Context) error {
	return nil
}

func (c CommandHandler) DeleteFeature(ctx context.Context) error {
	return nil
}
