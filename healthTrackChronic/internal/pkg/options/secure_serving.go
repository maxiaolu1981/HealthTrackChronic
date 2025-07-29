package options

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
