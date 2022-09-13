package feature

type GetFeatureQuery struct {
	Name string
}

type GetFeatureQueryResponse struct {
	Name       string
	MinVersion string
	Coverage   float32
}

type GetActiveFeaturesByUserIdQuery struct {
	UserId  uint
	Version string
}

type GetActiveFeaturesByUserIdQueryResponse struct {
	FeaturesNames []FeatureName
}

type FeatureName struct {
	Name string
}
