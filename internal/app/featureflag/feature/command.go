package feature

type UpsertFeatureCommand struct {
	Name       string
	MinVersion string
	Coverage   float32
}

type DeleteFeatureCommand struct {
	Name string
}
