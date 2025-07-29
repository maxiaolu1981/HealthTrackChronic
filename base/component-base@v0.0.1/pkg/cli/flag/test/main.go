package main

import (
	"log"
	"os"

	"github.com/maxiaolu1981/component-base/pkg/cli/flag"
	"github.com/spf13/pflag"
)

func main() {
	// 创建一个pflag标志集
	flags := pflag.NewFlagSet("myapp", pflag.ExitOnError)

	// 定义自定义标志（使用连字符风格）
	flags.String("log-level", "info", "日志级别")
	flags.Int("port", 8080, "服务端口")

	// 初始化标志集（标准化+整合标准flag）
	flag.InitFlags(flags)
	// 解析命令行参数
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("解析标志失败: %v", err)
	}
	// 打印所有标志信息（调试用）
	flag.PrintFlags(flags)
}

/*
# 连字符风格（推荐）
myapp --log-level debug --port 8081

# 下划线风格（会被自动转换，不推荐但兼容）
myapp --log_level debug --port 8081

*/
