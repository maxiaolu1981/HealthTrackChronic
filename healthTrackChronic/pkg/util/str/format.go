// 字符串相关辅助函数（格式化、截取、拼接、验证等）
package str

import (
	"runtime"
	"strings"
)

// 针对 Windows 系统的可执行文件，
// 去除其 .exe 后缀并转为小写，其他系统则保持原名称不变，
// 最终返回统一的 “基础名称”。比如在yalm配置文件中,保持一致
func FormatBaseName(baseName string) string {
	if runtime.GOOS == "windows" {
		baseName = strings.ToLower(baseName)
		baseName = strings.TrimSuffix(baseName, ".exe")
	}
	return baseName
}
