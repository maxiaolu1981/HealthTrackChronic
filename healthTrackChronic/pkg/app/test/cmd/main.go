package main

import (
	"fmt"
	"os"

	"github.com/maxiaolu1981/healthTrackChronic/pkg/app"
)

// 实现CliOptions接口的测试结构体
type TestOptions struct {
	Flag1 string
	Flag2 int
}

func (t *TestOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("test")
	fs.StringVar(&t.Flag1, "flag1", "default", "test flag1")
	fs.IntVar(&t.Flag2, "flag2", 0, "test flag2")
	return
}

func (t *TestOptions) Validate() []error {
	var errors []error
	if t.Flag1 == "" {
		errors = append(errors, fmt.Errorf("flag1 cannot be empty"))
	}
	if t.Flag2 < 0 {
		errors = append(errors, fmt.Errorf("flag2 must be non-negative"))
	}
	return errors
}

func main() {
	cmd1 := app.NewCommand(
		"cmd1",             // 命令名称（在终端中输入的指令）
		"cmd1 description", // 命令描述（--help时显示）
		app.WithCommandOptions(&TestOptions{}),
		app.WithCommandRunFunc(func(args []string) error {
			// 命令执行逻辑
			fmt.Println("cmd1 执行成功！")
			if len(args) > 0 {
				fmt.Printf("收到参数: %v\n", args)
			}
			return nil
		}),
	)
	// 2. 可选：添加子命令
	subCmd := app.NewCommand(
		"sub", // 子命令名称
		"sub command description",
		app.WithCommandRunFunc(func(args []string) error {
			fmt.Println("子命令 sub 执行成功！")
			return nil
		}),
	)
	cmd1.AddCommand(subCmd) // 将子命令添加到cmd1

	// 3. 转换为Cobra命令并执行
	cobraCmd := cmd1.CobraComand()             // 修正拼写错误：CobraComand -> CobraCommand
	if err := cobraCmd.Execute(); err != nil { // 执行Cobra命令
		fmt.Printf("命令执行失败: %v\n", err)
		os.Exit(1)
	}

}
