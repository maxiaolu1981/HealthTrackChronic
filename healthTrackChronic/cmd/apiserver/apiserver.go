package main

import (
	"fmt"
	"os"

	"github.com/maxiaolu1981/healthTrackChronic/pkg/app"

	cliflag "github.com/maxiaolu1981/component-base/pkg/cli/flag"
)

func main() {
	/*
		//用于生成一个基于当前时间的唯一随机数种子
		rand.Seed(time.Now().UTC().UnixNano())
		//设置最多能同时使用的 CPU 核心数
		if os.Getenv("GOMAXPROCS") == "" {
			runtime.GOMAXPROCS(runtime.NumCPU())
		}
		fmt.Println(runtime.GOMAXPROCS(0))
	*/

	

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
