package entity

type User struct {
	Id      uint
	Version Version
}

func NewUser(id uint, version Version) (User, error) {
	return User{id, version}, nil
}

//GetFeatureFlags checks the activation state of a feature for a specific user based on current features in
// the system and previous evaluated features for the user. Worst case time complexity is T = 2 * O(n) = O(n) where n is
// the  number of features in the system. Disclaimer: This function is not Clean at All! =))
func (u User) GetFeatureFlags(previousFeatureFlags []*FeatureWithFlag,
	allFeatures []*Feature) []*FeatureWithFlag {
	features := make([]*FeatureWithFlag, 0, len(allFeatures))
	oldFeaturesMap := createFeaturesMap(previousFeatureFlags)

	for _, feature := range allFeatures {
		if feature.Rule.HasMinVersion() &&
			u.Version.compareTo(feature.Rule.MinVersion) == -1 {
			features = append(features, &FeatureWithFlag{
				Name:     feature.Name,
				IsActive: false,
			})
			continue
		}

		if feature.Rule.HasCoverage() {
			oldFeature, exists := oldFeaturesMap[feature.Name]
			if exists && !oldFeature.isActive {
				if oldFeature.coverage < feature.Rule.Coverage {
					if feature.isCovered() {
						features = append(features, &FeatureWithFlag{
							Name:     feature.Name,
							IsActive: true,
							Coverage: feature.Rule.Coverage,
						})
						continue
					}
				}
				features = append(features, &FeatureWithFlag{
					Name:     feature.Name,
					IsActive: false,
					Coverage: feature.Rule.Coverage,
				})
				continue
			}

			if exists && oldFeature.isActive {
				features = append(features, &FeatureWithFlag{
					Name:     feature.Name,
					IsActive: true,
					Coverage: feature.Rule.Coverage,
				})
				continue
			}

			if feature.isCovered() {
				features = append(features, &FeatureWithFlag{
					Name:     feature.Name,
					IsActive: true,
					Coverage: feature.Rule.Coverage,
				})
				continue
			}

			features = append(features, &FeatureWithFlag{
				Name:     feature.Name,
				IsActive: false,
				Coverage: feature.Rule.Coverage,
			})
			continue
		}

		features = append(features, &FeatureWithFlag{
			Name:     feature.Name,
			IsActive: true,
		})

	}

	return features
}

func createFeaturesMap(features []*FeatureWithFlag) map[string]featureDTO {
	featureMap := make(map[string]featureDTO)
	for _, feature := range features {
		featureMap[feature.Name] = featureDTO{
			coverage: feature.Coverage,
			isActive: feature.IsActive,
		}
	}

	return featureMap
}

type featureDTO struct {
	coverage float32
	isActive bool
}
