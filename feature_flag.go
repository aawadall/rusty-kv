package main

// FeatureFlag - feature flag
type FeatureFlag struct {
	// Name - name of the feature flag
	Name string
	// Description - description of the feature flag
	Description string
	// Enabled - whether the feature flag is enabled
	Enabled bool
}

// FeatureFlagManager - feature flag manager
type FeatureFlagManager struct {
	// FeatureFlags - list of feature flags
	FeatureFlags []FeatureFlag
}

// NewFeatureFlagManager - create a new feature flag manager
func NewFeatureFlagManager() *FeatureFlagManager {
	return &FeatureFlagManager{
		FeatureFlags: []FeatureFlag{},
	}
}

// Add - add a feature flag to the manager
func (ffm *FeatureFlagManager) Add(name string, description string, enabled bool) {
	ffm.FeatureFlags = append(ffm.FeatureFlags, FeatureFlag{
		Name:        name,
		Description: description,
		Enabled:     enabled,
	})
}

// Remove - remove a feature flag from the manager
func (ffm *FeatureFlagManager) Remove(name string) {
	for i, featureFlag := range ffm.FeatureFlags {
		if featureFlag.Name == name {
			ffm.FeatureFlags = append(ffm.FeatureFlags[:i], ffm.FeatureFlags[i+1:]...)
		}
	}
}

// Get
func (ffm *FeatureFlagManager) Get(name string) *FeatureFlag {
	for _, featureFlag := range ffm.FeatureFlags {
		if featureFlag.Name == name {
			return &featureFlag
		}
	}
	return nil
}

// GetFeatureFlags - get all feature flags
func (ffm *FeatureFlagManager) GetFeatureFlags() []FeatureFlag {
	return ffm.FeatureFlags
}

// IsEnabled - check if a feature flag is enabled
func (ffm *FeatureFlagManager) IsEnabled(name string) bool {
	featureFlag := ffm.Get(name)
	if featureFlag != nil {
		return featureFlag.Enabled
	}
	return false
}

// Enable
func (ffm *FeatureFlagManager) Enable(name string) {
	featureFlag := ffm.Get(name)
	if featureFlag != nil {
		featureFlag.Enabled = true
	}
}

// Disable
func (ffm *FeatureFlagManager) Disable(name string) {
	featureFlag := ffm.Get(name)
	if featureFlag != nil {
		featureFlag.Enabled = false
	}
}
