package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {

	//用于生成一个基于当前时间的唯一随机数种子
	rand.Seed(time.Now().UTC().UnixNano())
	//设置最多能同时使用的 CPU 核心数
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	fmt.Println(runtime.GOMAXPROCS(0))

}
