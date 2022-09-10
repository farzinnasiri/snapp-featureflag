package feature

import "fmt"

type Feature struct {
	Name   string
	Policy Policy
}

func NewFeature(name string, policy Policy) (Feature, error) {
	//todo validate
	return Feature{name, policy}, nil
}

type Policy struct {
	Coverage   float32
	MinVersion *Version
}

func NewPolicy(coverage float32, minVersion *Version) (Policy, error) {
	//todo validate
	return Policy{coverage, minVersion}, nil
}

type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func NewVersion(major uint, minor uint, patch uint) (Version, error) {
	return Version{major, minor, patch}, nil
}

func (v Version) toString() string {
	return fmt.Sprintf(
		"%d.%d.%d",
		v.Major, v.Minor, v.Patch)
}

func (v Version) compareTo(v1 Version) int {
	if v.Major > v1.Major {
		return 1
	} else if v.Major < v1.Major {
		return -1
	} else {
		if v.Minor > v1.Minor {
			return 1
		} else if v.Minor < v1.Minor {
			return -1
		} else {
			if v.Patch > v1.Patch {
				return 1
			} else if v.Patch < v1.Patch {
				return -1
			} else {
				return 0
			}
		}
	}
}
