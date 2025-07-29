package app

/*
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

*/

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/fatih/color"
	cliflag "github.com/maxiaolu1981/component-base/pkg/cli/flag"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var Opts = &TestOptions{}

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

// 测试NewCommand函数
func TestNewCommand(t *testing.T) {
	cmd := NewCommand("test", "test description")
	assert.Equal(t, "test", cmd.usage)
	assert.Equal(t, "test description", cmd.desc)
	assert.Nil(t, cmd.options)
	assert.Nil(t, cmd.runFunc)
}

// 测试WithCommandOptions
func TestWithCommandOptions(t *testing.T) {
	opts := &TestOptions{}
	cmd := NewCommand("test", "desc", WithCommandOptions(opts))
	assert.Equal(t, opts, cmd.options)
}

// 测试WithCommandRunFunc
func TestWithCommandRunFunc(t *testing.T) {
	var executed bool
	runFunc := func(args []string) error {
		executed = true
		t.Log("我是回调函数.........")
		return nil
	}

	cmd := NewCommand("test", "desc", WithCommandRunFunc(runFunc))

	// 调用runCommand方法，触发函数执行
	cobraCmd := &cobra.Command{}
	cmd.runCommand(cobraCmd, []string{})

	// 验证函数是否被执行
	assert.True(t, executed, "RunCommandFunc should be executed")
}

// 测试AddCommand和AddCommands
func TestAddCommand(t *testing.T) {
	parent := NewCommand("parent", "parent desc")
	child1 := NewCommand("child1", "child1 desc")
	child2 := NewCommand("child2", "child2 desc")

	parent.AddCommand(child1)
	parent.AddCommands(child2)

	assert.Len(t, parent.commands, 2)
	assert.Contains(t, parent.commands, child1)
	assert.Contains(t, parent.commands, child2)
}

// 测试cobraCommand
func TestCobraCommand(t *testing.T) {
	// 测试无options和runFunc的情况
	cmd1 := NewCommand("cmd1", "cmd1 desc")
	cobraCmd1 := cmd1.CobraComand()
	assert.Equal(t, "cmd1", cobraCmd1.Use)
	assert.Equal(t, "cmd1 desc", cobraCmd1.Short)
	assert.Nil(t, cobraCmd1.Run)
	assert.Len(t, cobraCmd1.Commands(), 0)

	// 测试有子命令的情况
	child := NewCommand("child", "child desc")
	cmd1.AddCommand(child)
	cobraCmd1 = cmd1.CobraComand()
	assert.Len(t, cobraCmd1.Commands(), 1)

	// 测试有options的情况
	opts := &TestOptions{}
	cmd2 := NewCommand("cmd2", "cmd2 desc", WithCommandOptions(opts))
	cobraCmd2 := cmd2.CobraComand()
	assert.NotNil(t, cobraCmd2.Flags().Lookup("flag1"))
	assert.NotNil(t, cobraCmd2.Flags().Lookup("flag2"))

	// 测试有runFunc的情况
	runFuncCalled := false
	runFunc := func(args []string) error {
		runFuncCalled = true
		return nil
	}
	cmd3 := NewCommand("cmd3", "cmd3 desc", WithCommandRunFunc(runFunc))
	cobraCmd3 := cmd3.CobraComand()
	cobraCmd3.SetArgs([]string{})
	cobraCmd3.Execute()
	assert.True(t, runFuncCalled)
}

// 测试runCommand
func TestRunCommand(t *testing.T) {
	// 捕获标准输出
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
		w.Close()
	}()

	// 测试错误处理
	runFunc := func(args []string) error {
		return fmt.Errorf("test error")
	}
	cmd := NewCommand("test", "desc", WithCommandRunFunc(runFunc))
	cobraCmd := &cobra.Command{}
	cmd.runCommand(cobraCmd, []string{})

	// 读取输出
	w.Close()
	out, _ := io.ReadAll(r)

	expected := color.RedString("Error:") + " test error\n"
	assert.Contains(t, string(out), expected)
}

// 测试App的AddCommand和AddCommands
func TestAppAddCommand(t *testing.T) {
	app := &App{}
	cmd1 := NewCommand("cmd1", "desc1")
	cmd2 := NewCommand("cmd2", "desc2")

	app.AddCommand(cmd1)
	app.AddCommands(cmd2)

	assert.Len(t, app.commands, 2)
	assert.Contains(t, app.commands, cmd1)
	assert.Contains(t, app.commands, cmd2)
}

// 测试帮助命令标志
func TestHelpCommandFlag(t *testing.T) {
	cmd := NewCommand("test", "desc")
	cobraCmd := cmd.CobraComand()
	assert.NotNil(t, cobraCmd.Flags().Lookup("help"))
}

// 测试标志排序
func TestFlagSorting(t *testing.T) {
	cmd := NewCommand("test", "desc")
	cobraCmd := cmd.CobraComand()
	assert.False(t, cobraCmd.Flags().SortFlags)
}

// 测试输出设置
func TestOutputSetting(t *testing.T) {
	cmd := NewCommand("test", "desc")
	cobraCmd := cmd.CobraComand()
	assert.Equal(t, os.Stdout, cobraCmd.OutOrStdout())
}
