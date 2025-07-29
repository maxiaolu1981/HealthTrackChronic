package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// RecommendedHomeDir defines the default directory used to place all iam service configurations.
	RecommendedHomeDir = ".iam"

	// RecommendedEnvPrefix defines the ENV prefix used by all iam service.
	RecommendedEnvPrefix = "IAM"
)

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		Jwt: &JwtInfo{
			Realm:     "healthTrackChronic jwt",
			Timeout:   1 * time.Hour,
			MaxRefuse: 1 * time.Hour,
		},
	}
}

type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey //证书和私钥
}

// 存储 TLS 证书和私钥的文件路径配
type CertKey struct {
	//证书文件路径，包含 PEM 格式的证书（可能包含完整证书链）
	CretFile string
	//私钥文件路径，包含 PEM 格式的私钥（对应 CertFile 中的证书）。
	KeyFile string
}

type InsecureServingInfo struct {
	Address string
}

type JwtInfo struct {
	//认证域名称，用于 WWW-Authenticate 响应头（如 Bearer realm="iam jwt"）
	Realm string //认证域（Realm）：用于提示用户认证的范围或领域。
	//签名和验证 JWT 的密钥（建议使用至少 32 字节的随机字符串）
	Key string //签名密钥（Key）：用于生成和验证 JWT 的密钥。
	//JWT 的有效时长（如 1h 表示 1 小时）。超过此时长后，令牌将失效。
	Timeout time.Duration //令牌有效期（Timeout）：控制 JWT 的有效时长，防止令牌被长期滥用。
	//令牌可刷新的最大时间窗口（如 24h 表示 24 小时内可刷新）。
	MaxRefuse time.Duration //最大刷新时间（MaxRefresh）：允许刷新令牌的最长时间窗口。
}
