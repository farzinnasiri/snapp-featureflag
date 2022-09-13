package entity

import (
	"errors"
	"math/rand"
)

type Feature struct {
	Name string
	Rule Rule
}

type CreateFeatureParams struct {
	Name       string
	Coverage   float32
	MinVersion string
}

type FeatureNameWithActivation struct {
	Name     string
	IsActive bool
}

func NewFeatureFromParams(params CreateFeatureParams) (*Feature, error) {
	minVersion, err := getVersion(params.MinVersion)
	if err != nil {
		return nil, err
	}
	policy, err := getPolicy(params.Coverage, minVersion)
	if err != nil {
		return nil, err
	}

	feature, err := NewFeature(params.Name, *policy)
	if err != nil {
		return nil, err
	}

	return feature, nil
}

func NewFeature(name string, policy Rule) (*Feature, error) {
	feature := Feature{name, policy}
	if err := feature.validate(); err != nil {
		return nil, err
	}
	return &feature, nil
}

func getVersion(version string) (*Version, error) {
	if version != "" {
		minVersion, err := NewVersionFromString(version)
		if err != nil {
			return nil, err
		}
		return minVersion, minVersion.validate()
	} else {
		return nil, nil
	}
}

func getPolicy(coverage float32, minVersion *Version) (*Rule, error) {
	policy, err := NewRule(coverage, minVersion)
	if err != nil {
		return nil, err
	}
	if err = policy.validate(); err != nil {
		return nil, err
	}
	return policy, nil
}

func (f *Feature) validate() error {
	if f.Name == "" {
		return errors.New("name can not be empty")
	}
	return nil
}

func (f *Feature) Update(newFeature *Feature) error {
	if f.Rule.Coverage > newFeature.Rule.Coverage {
		return errors.New("can not reduce coverage")
	}
	f.Rule = newFeature.Rule

	return nil
}

func (f *Feature) isCovered() bool {
	coverage := int(f.Rule.Coverage * 100)

	r := rand.Intn(100)
	if coverage >= r {
		return true
	}
	return false

}
