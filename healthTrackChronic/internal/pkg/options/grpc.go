package options

import "github.com/spf13/pflag"

type GRPCOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
	MaxMsgSize  int    `json:"max-msg-size" mapstructure:"max-msg-size"`
}

func NewGRPCOptions() *GRPCOptions {
	return &GRPCOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8081,
		MaxMsgSize:  4 * 1024 * 1024,
	}
}

func (s *GRPCOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "grpc.bind-address", s.BindAddress, "用于监听 --grpc.bind-port 端口的 IP 地址（若要监听所有 IPv4 接口，设为 0.0.0.0；若要监听所有 IPv6 接口，设为 ::）。")
	fs.IntVar(&s.BindPort, "grpc.bind-port", s.BindPort, "用于提供未加密、未认证的 gRPC 访问的端口。默认假设已配置防火墙规则，确保该端口仅能从部署机器内部访问，且公网地址的 443 端口会代理到该端口（默认配置中由 nginx 实现此代理）。设置为 0 可禁用该端口。")
	fs.IntVar(&s.MaxMsgSize, "grpc.max-msg-size", s.MaxMsgSize, "gRPC最大消息长度")
}
