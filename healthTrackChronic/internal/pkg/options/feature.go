package options

import "github.com/maxiaolu1981/healthTrackChronic/internal/pkg/server"

type FeatureOptions struct {
	EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

func NewFeatureOptions() *FeatureOptions {
	defaults := server.NewConfig()
	return &FeatureOptions{
		EnableMetrics:   defaults.EnableMetrics,
		EnableProfiling: defaults.EnableProfiling,
	}
}
