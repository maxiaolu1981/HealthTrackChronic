package options

import (
	cliflag "github.com/maxiaolu1981/base/component-base/pkg/cli/flag"
	genericoptions "github.com/maxiaolu1981/healthTrackChronic/internal/pkg/options"
	"github.com/maxiaolu1981/healthTrackChronic/pkg/log"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	GRPCOptions             *genericoptions.GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"    mapstructure:"redis"`
	JwtOptions              *genericoptions.JwtOptions             `json:"jwt"      mapstructure:"jwt"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

func NewOptions() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		GRPCOptions:             genericoptions.NewGRPCOptions(),
		InsecureServing:         genericoptions.NewNewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		RedisOptions:            genericoptions.NewRedisOptions(),
		JwtOptions:              genericoptions.NewJwtOptions(),
		Log:                     log.NewOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
	}
}

func (opt Options) Flags() (fss cliflag.NamedFlagSets) {
	opt.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	opt.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	opt.JwtOptions.AddFlags(fss.FlagSet("jwt"))
	opt.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	opt.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	opt.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	opt.RedisOptions.AddFlags(fss.FlagSet("redis"))
	opt.Log.AddFlags(fss.FlagSet("log"))
	opt.FeatureOptions.AddFlags(fss.FlagSet("feature"))

	return fss
}
