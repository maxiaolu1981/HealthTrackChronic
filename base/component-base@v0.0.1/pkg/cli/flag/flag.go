/*
pflag（增强版命令行标志库，兼容标准库flag）的工具函数封装，
主要用于标准化标志（flag）的命名格式、初始化标志集以及打印
标志信息。
--标志名称的标准化（处理下划线_和连字符-的兼容）
--初始化标志集（整合标准库flag和pflag）
--打印标志信息（便于调试和日志记录）
*/
package flag

import (
	goflag "flag"
	"strings"

	"github.com/marmotedu/log"
	"github.com/spf13/pflag"
)

// 将标志名称中的下划线_自动转换为连字符-，实现两种命名风格的兼容
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// 作用：功能与WordSepNormalizeFunc类似（转换_为-），但会额外打印警告日志，提示用户下划线格式已过时。
// 使用场景：用于版本迁移（如从旧版本的下划线风格标志，逐步过渡到连字符风格），通过警告引导用户使用新格式。
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")
		log.Warnf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, nname)

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// 标准化标志集并整合标准库flag，确保标志解析的一致性。
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

/*
作用：遍历标志集中的所有标志，以DEBUG级别日志打印标志名称和对应的值。
实现逻辑：
通过flags.VisitAll遍历所有标志（包括默认值未被修改的标志）。
对每个标志，打印格式为FLAG: --<name>=<value>（如FLAG: --log-level="info"）。
用途：
调试时确认标志是否被正确解析（例如检查配置文件或命令行输入的标志是否生效）。
记录运行时的标志状态，便于问题排查。
*/
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debugf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
