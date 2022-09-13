package entity

type User struct {
	Id      uint
	Version Version
}

func NewUser(id uint, version Version) (User, error) {
	return User{id, version}, nil
}

//GetFeaturesActivationStates finds active features list in T = O(n) + O(m) time, where n is the length of all features
// and m is the length of previous active features. Worst case T = 2*O(n) = O(n)
func (u User) GetFeaturesActivationStates(previousActivatedFeatures []*FeatureWithFlag,
	allFeatures []*Feature) []*FeatureWithFlag {
	features := make([]*FeatureWithFlag, 0, len(allFeatures))
	oldFeaturesMap := createFeaturesMap(previousActivatedFeatures)

	for _, feature := range allFeatures {
		if feature.Rule.HasMinVersion() &&
			u.Version.compareTo(feature.Rule.MinVersion) == -1 {
			continue
		}

		if feature.Rule.HasCoverage() {
			wasActive, exists := oldFeaturesMap[feature.Name]
			if exists && !wasActive {
				features = append(features, &FeatureWithFlag{
					Name:     feature.Name,
					IsActive: false,
				})
				continue
			}
			if (exists && wasActive) || (feature.isCovered()) {
				features = append(features, &FeatureWithFlag{
					Name:     feature.Name,
					IsActive: true,
				})
			}
			features = append(features, &FeatureWithFlag{
				Name:     feature.Name,
				IsActive: false,
			})
			continue
		}

		features = append(features, &FeatureWithFlag{
			Name:     feature.Name,
			IsActive: false,
		})

	}

	return features
}

func createFeaturesMap(features []*FeatureWithFlag) map[string]bool {
	featureMap := make(map[string]bool)
	for _, feature := range features {
		featureMap[feature.Name] = feature.IsActive
	}

	return featureMap

}
