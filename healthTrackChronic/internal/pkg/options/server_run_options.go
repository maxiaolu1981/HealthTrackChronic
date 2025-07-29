package options

import (
	"github.com/spf13/pflag"

	"github.com/maxiaolu1981/healthTrackChronic/internal/pkg/server"
)

type ServerRunOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Health      bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()
	return &ServerRunOptions{
		Mode:        defaults.Mode,
		Health:      defaults.Healthz,
		Middlewares: defaults.Middlewares,
	}
}

func (s *ServerRunOptions) AddFlags(fs *Pflag.FlagSet)
