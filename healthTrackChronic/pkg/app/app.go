package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RunFunc func(baseName string)
type Option func(*App)

type App struct {
	baseName    string //二进制文件名（如"myapp"）
	name        string //应用名称（如"我的命令行工具"）
	description string
	options     CliOptions           // 配置选项（用于解析命令行参数和配置文件
	runFunc     RunFunc              // 根命令的执行逻辑
	silence     bool                 // 是否静默模式（不打印启动信息）
	noVersion   bool                 // 是否禁用版本标志（--version）
	noConfig    bool                 // 是否禁用配置文件支持
	command     []*Command           // 子命令集合
	args        cobra.PositionalArgs // 位置参数验证函数
	cmd         *cobra.Command       // 内部封装的Cobra命令
}

// 绑定配置选项（用于解析命令行参数和配置文件）。
func withOptions(cli CliOptions) Option {
	return func(app *App) {
		app.options = cli
	}
}

// 设置应用启动后的核心业务逻辑函数（runFunc）
func withRunFunc(run RunFunc) Option {
	return func(app *App) {
		app.runFunc = run
	}
}

// 设置应用的详细描述（显示在--help中）
func withDescription(desc string) Option {
	return func(app *App) {
		app.description = desc
	}
}

// 启用静默模式（不打印启动过程信息）。
func withSilence() Option {
	return func(app *App) {
		app.silence = true
	}
}

// 禁用版本查询（--version标志无效）
func withNoVersion() Option {
	return func(app *App) {
		app.noVersion = true
	}
}

// 禁用配置文件支持（不解析--config指定的文件）。
func withNoConfig() Option {
	return func(app *App) {
		app.noConfig = true
	}
}

// 自定义位置参数验证规则（如允许特定参数）
func WithValidArgs(posit cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = posit
	}
}

// 默认位置参数规则（禁止任何位置参数，避免误输入）
func WithDefaultValidArgs() Option {
	return func(app *App) {
		app.args = func(cmd *cobra.Command, args []string) error {
			for _, o := range args {
				if len(o) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}

}

// 作用：初始化App结构体，通过Option函数配置属性，最终调用buildCommand构建 Cobra 命令。
// 参数：name（应用名称）、basename（二进制文件名）、opts（配置选项）。
func NewApp(name, baseName string, opts ...Option) *App {
	app := &App{
		name:     name,
		baseName: baseName,
	}
	for _, o := range opts {
		o(app)
	}

	return app
}

/*
将App的配置转换为cobra.Command实例（a.cmd），是连接App与 Cobra 底层的核心方法，步骤如下：
初始化cobra.Command，设置基本信息（用法、描述、输入输出流等）。
添加子命令：将a.commands中的子命令转换为 Cobra 命令并关联到根命令。
处理标志（flags）：
添加应用自定义标志（通过a.options.Flags()获取）。
添加全局标志：版本（--version，通过verflag）、配置文件（--config）等。
绑定命令执行函数：将a.runCommand设为 Cobra 命令的RunE（执行入口）。
*/
func buildCommand(app *App) {
	cmd := &cobra.Command{
		Use: app.baseName,
	}
}
