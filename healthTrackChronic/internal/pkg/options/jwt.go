package options

import (
	"time"

	"github.com/maxiaolu1981/healthTrackChronic/internal/pkg/server"
	"github.com/spf13/pflag"
	"github.com/asaskevich/govalidator"
)

type JwtOptions struct {
	Realm      string        `json:"realm"       mapstructure:"realm"`
	Key        string        `json:"key"         mapstructure:"key"`
	Timeout    time.Duration `json:"timeout"     mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

func NewJwtOptions() *JwtOptions {
	defaults := server.NewConfig()
	return &JwtOptions{
		Realm:      defaults.Jwt.Realm,
		Key:        defaults.Jwt.Key,
		Timeout:    defaults.Jwt.Timeout,
		MaxRefresh: defaults.Jwt.MaxRefuse,
	}
}

func (s *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}
	fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "向用户显示的认证域名称。")
	fs.StringVar(&s.Key, "jwt.ket", s.Key, "用于签署 JWT 令牌的私钥")
	fs.DurationVar(&s.Timeout, "jwt.timeout", s.Timeout, "WT token.超时时间")
	fs.DurationVar(&s.MaxRefresh, "jwt.max-refresh", s.MaxRefresh, "允许客户端在 MaxRefresh 时间窗口内刷新其令牌。补充说明：")
}

func (s *JwtOptions) Validate() []error {
	var errs []error

	if !govalidator.StringLength(s.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("--secret-key must larger than 5 and little than 33"))
	}

	return errs
}
