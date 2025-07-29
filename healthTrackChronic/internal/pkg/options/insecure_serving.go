package options

import "github.com/spf13/pflag"

type InsecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewNewInsecureServingOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
	}
}

func (s *InsecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "insecure.bind-address", s.BindAddress, ""+
		"用于监听 --insecure.bind-port 端口的 IP 地址。 "+
		"若要监听所有 IPv4 接口，设为 0.0.0.0；若要监听所有 IPv6 接口，设为 ::）。")
	fs.IntVar(&s.BindPort, "insecure.bind-port", s.BindPort, "用于提供未加密、未认证访问的端口。默认假设已配置防火墙规则，确保该端口无法从部署机器外部访问，且 IAM 公网地址的 443 端口会代理到该端口（默认配置中由 nginx 实现此代理）。设置为 0 可禁用该端口。")
}
