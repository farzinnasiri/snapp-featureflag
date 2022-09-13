package entity

type User struct {
	Id      uint
	Version Version
}

func NewUser(id uint, version Version) (User, error) {
	return User{id, version}, nil
}

//GetActiveFeatures finds active features list in T = O(n) + O(m) time, where n is the length of all features
// and m is the length of previous active features. Worst case T = 2*O(n) = O(n)
func (u User) GetActiveFeatures(previousActivatedFeatures []*Feature,
	allFeatures []*Feature) []*Feature {
	activeFeatures := make([]*Feature, 0, len(allFeatures))
	oldFeaturesMap := createFeaturesMap(previousActivatedFeatures)

	for _, feature := range allFeatures {
		if feature.Rule.HasMinVersion() &&
			u.Version.compareTo(feature.Rule.MinVersion) == -1 {
			continue
		}

		if feature.Rule.HasCoverage() {
			_, exists := oldFeaturesMap[feature.Name]
			if exists || feature.isCovered() {
				activeFeatures = append(activeFeatures, feature)
			}
			continue
		}

		activeFeatures = append(activeFeatures, feature)

	}

	return activeFeatures
}

func createFeaturesMap(features []*Feature) map[string]byte {
	featureMap := make(map[string]byte)
	for _, feature := range features {
		featureMap[feature.Name] = 1
	}

	return featureMap

}
