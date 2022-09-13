package entity

import "errors"

type Rule struct {
	Coverage   float32
	MinVersion *Version
}

func NewRule(coverage float32, minVersion *Version) (*Rule, error) {
	return &Rule{coverage, minVersion}, nil
}

func (r *Rule) validate() error {
	if r.Coverage < 0 {
		return errors.New("r coverage should be greater than zero")
	}
	if r.Coverage == 0 && r.MinVersion == nil {
		return errors.New("policy can not be nil")
	}
	if r.MinVersion != nil {
		return r.MinVersion.validate()
	}
	return nil

}

func (r *Rule) HasCoverage() bool {
	return r.Coverage != 0
}

func (r *Rule) HasMinVersion() bool {
	return r.MinVersion != nil
}
