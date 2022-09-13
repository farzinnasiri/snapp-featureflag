package feature

type CreateFeatureCommand struct {
	Name       string
	MinVersion string
	Coverage   float32
}

type UpdateFeatureCommand struct {
	Name       string
	MinVersion string
	Coverage   float32
}

type DeleteFeatureCommand struct {
	Name string
}
