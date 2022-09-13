package entity

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserEntityTestSuite struct {
	suite.Suite
}

func TestUserEntity(t *testing.T) {
	suite.Run(t, new(UserEntityTestSuite))
}

func (t *UserEntityTestSuite) SetupSuite() {
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsUserVersionIsAboveMinVersionActivateFeature() {
	user, _ := NewUser(1, Version{1, 0, 0})

	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{MinVersion: &Version{0, 5, 0}},
	}}

	featureFlags := user.GetFeatureFlags(make([]*FeatureWithFlag, 0), allFeatures)

	t.Require().Len(featureFlags, 1)
	t.Require().Equal("test", featureFlags[0].Name)

}

func (t *UserEntityTestSuite) TestGetFeatureFlagsPartialFeatureIsActiveForEnoughUsers() {

	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.5},
	}}

	var count float64
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{1, 0, 0})
		featureFlags := user.GetFeatureFlags(make([]*FeatureWithFlag, 0), allFeatures)
		t.Require().Len(featureFlags, 1)

		if featureFlags[0].IsActive {
			count++
		}
	}
	t.Require().Greater(count/1000, 0.45)
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsCombinedFeatureIsActiveForEnoughUsers() {
	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.5, MinVersion: &Version{0, 5, 0}},
	}}

	var count float64
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{1, 0, 0})
		featureFlags := user.GetFeatureFlags(make([]*FeatureWithFlag, 0), allFeatures)
		t.Require().Len(featureFlags, 1)
		if featureFlags[0].IsActive {
			count++
		}
	}
	t.Require().Greater(count/1000, 0.45)
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsCombinedFeatureIsNotActiveForOldUsers() {
	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.5, MinVersion: &Version{0, 5, 0}},
	}}

	var count float64
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{0, 4, 5})
		featureFlags := user.GetFeatureFlags(make([]*FeatureWithFlag, 0), allFeatures)
		t.Require().Len(featureFlags, 1)
		if featureFlags[0].IsActive {
			count++
		}
	}
	t.Require().Zero(count)
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsIfGlobalFeatureActiveForAllUsers() {
	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 1},
	}}

	var count int
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{0, 4, 5})
		featureFlags := user.GetFeatureFlags(make([]*FeatureWithFlag, 0), allFeatures)
		t.Require().Len(featureFlags, 1)
		if featureFlags[0].IsActive {
			count++
		}
	}
	t.Require().Equal(1000, count)

}

func (t *UserEntityTestSuite) TestGetFeatureFlagsIfPartialFeatureAlreadyActiveRemainActive() {
	user, _ := NewUser(1, Version{1, 0, 0})

	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.01},
	}}

	previousFeatureFlags := []*FeatureWithFlag{&FeatureWithFlag{
		Name:     "test",
		IsActive: true,
	}}

	featureFlags := user.GetFeatureFlags(previousFeatureFlags, allFeatures)

	t.Require().Len(featureFlags, 1)
	t.Require().Equal(true, featureFlags[0].IsActive)
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsIfPartialFeatureNotActiveAndCoverageNotChangeRemainNotActive() {

	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.5},
	}}

	previousFeatureFlags := []*FeatureWithFlag{&FeatureWithFlag{
		Name:     "test",
		IsActive: false,
		Coverage: 0.5,
	}}

	var count int
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{0, 4, 5})
		featureFlags := user.GetFeatureFlags(previousFeatureFlags, allFeatures)
		t.Require().Len(featureFlags, 1)
		if featureFlags[0].IsActive {
			count++
		}
	}

	t.Require().Zero(count)
}

func (t *UserEntityTestSuite) TestGetFeatureFlagsIfPartialFeatureNotActiveAndCoverageChangeActive() {
	allFeatures := []*Feature{&Feature{
		Name: "test",
		Rule: Rule{Coverage: 0.5},
	}}

	previousFeatureFlags := []*FeatureWithFlag{&FeatureWithFlag{
		Name:     "test",
		IsActive: false,
		Coverage: 0.3,
	}}

	var count float64
	for i := 0; i < 1000; i++ {
		user, _ := NewUser(1, Version{0, 4, 5})
		featureFlags := user.GetFeatureFlags(previousFeatureFlags, allFeatures)
		t.Require().Len(featureFlags, 1)
		if featureFlags[0].IsActive {
			count++
		}
	}

	t.Require().Greater(count/1000, 0.45)

}
