package options

import "github.com/spf13/pflag"

type SecureServingOptions struct {
	// 监听地址（如 "0.0.0.0"）
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	// 监听端口（如 443）
	BindPort int `json:"bind-port"    mapstructure:"bind-port"`
	// 是否必须绑定端口（若为 true，BindPort 不能为 0）
	Required bool
	// TLS 证书配置
	ServerCert GeneratableKeyCert `json:"tls"          mapstructure:"tls"`
	// AdvertiseAddress net.IP
}

// 支持自动生成的证书配置
type GeneratableKeyCert struct {
	// 手动指定证书文件（优先级高）
	CertKey CertKey `json:"cert-key" mapstructure:"cert-key"`
	// 自动生成证书的存储目录（如 "/etc/certs"）
	CertDirectory string `json:"cert-dir"  mapstructure:"cert-dir"`
	// 证书文件名前缀（如 "server"，生成 server.crt 和 server.key）
	PairName string `json:"pair-name" mapstructure:"pair-name"`
}

// 直接指定证书文件路径
type CertKey struct {
	// 证书文件路径（如 "/etc/certs/server.crt"）
	CertFile string `json:"cert-file"        mapstructure:"cert-file"`
	//私钥文件路径（如 "/etc/certs/server.key"）
	KeyFile string `json:"private-key-file" mapstructure:"private-key-file"`
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8443,
		Required:    true,
		ServerCert: GeneratableKeyCert{
			PairName:      "healthTrackChronic",
			CertDirectory: "/var/run/healthTrackChronic",
		},
	}
}

func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "secure.bind-address", s.BindAddress, "用于监听 --secure.bind-port 端口的 IP 地址。相关网络接口必须可被引擎的其他组件以及 CLI / 网页客户端访问。若留空，将使用所有接口（0.0.0.0 对应所有 IPv4 接口，:: 对应所有 IPv6 接口）。")
	desc := "用于提供带身份验证和授权的 HTTPS 服务的端口。"
	if s.Required {
		desc += " 该端口不能通过设置为 0 来关闭。"
	} else {
		desc += " 若设置为 0，则完全不提供 HTTPS 服务。"
	}
	fs.IntVar(&s.BindPort, "secure.bind-port", s.BindPort, desc)

	fs.StringVar(&s.ServerCert.CertDirectory, "secure.tls.cert-dir", s.ServerCert.CertDirectory, "若已提供 --secure.tls.cert-key.cert-file 和 --secure.tls.cert-key.private-key-file（即证书文件和私钥文件的具体路径），则此标志（指该目录配置）将被忽略。")

	fs.StringVar(&s.ServerCert.PairName, "secure.tls.pair-name", s.ServerCert.PairName, "与 --secure.tls.cert-dir 结合使用的证书和私钥文件名前缀。生成的证书文件名为 <证书目录>/<名称对>.crt，私钥文件名为 <证书目录>/<名称对>.key。")

	fs.StringVar(&s.ServerCert.CertKey.CertFile, "secure.tls.cert-key.cert-file", s.ServerCert.CertKey.CertFile, "用于 HTTPS 的默认 X.509 证书文件（若存在 CA 证书，需拼接在服务器证书之后）。")

	fs.StringVar(&s.ServerCert.CertKey.KeyFile, "secure.tls.cert-key.private-key-file",
		s.ServerCert.CertKey.KeyFile, ""+
			"与 --secure.tls.cert-key.cert-file 对应的默认 X.509 私钥文件。")
}
