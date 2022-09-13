package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func NewVersion(major uint, minor uint, patch uint) (*Version, error) {
	return &Version{major, minor, patch}, nil
}

func NewVersionFromString(versionStr string) (*Version, error) {
	versionParams := strings.Split(versionStr, ".")
	if len(versionParams) != 3 {
		return nil, errors.New(
			fmt.Sprintf(
				"expected 3 parameters in version, got %d",
				len(versionParams)))
	}
	major, err := strconv.ParseUint(versionParams[0], 0, 32)
	if err != nil {
		return nil, err
	}
	minor, err := strconv.ParseUint(versionParams[1], 0, 32)
	if err != nil {
		return nil, err
	}
	patch, err := strconv.ParseUint(versionParams[2], 0, 32)
	if err != nil {
		return nil, err
	}

	return NewVersion(uint(major), uint(minor), uint(patch))

}

func NewValidVersionFromString(versionStr string) (*Version, error) {
	version, err := NewVersionFromString(versionStr)
	if err != nil {
		return nil, err
	}
	if err = version.validate(); err != nil {
		return nil, err
	}
	return version, nil
}

func (v *Version) validate() error {
	if v.Major == 0 && v.Minor == 0 && v.Patch == 0 {
		return errors.New("all version fields can not be zero")
	}
	return nil
}

func (v *Version) ToString() string {
	return fmt.Sprintf(
		"%d.%d.%d",
		v.Major, v.Minor, v.Patch)
}

//compareTo returns 1 if the receiver is greater than the argument, 0 if their equal, and -1 otherwise
func (v *Version) compareTo(v1 *Version) int {
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
