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

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mode, "server.mode", s.Mode, "以指定的服务器模式启动服务器。支持的服务器模式：debug（调试）、test（测试）、release（发布）。")
	fs.BoolVar(&s.Health, "server.healthz", s.Health, "添加自身就绪性检查，并配置 /healthz 路由")
	fs.StringSliceVar(&s.Middlewares, "server.middlewares", s.Middlewares, "服务器允许使用的中间件列表，以逗号分隔。若此列表为空，则使用默认中间件")
	
}
