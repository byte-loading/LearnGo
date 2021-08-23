package main

import (
	"fmt"
	"os"
)

// 获取用户命令行参数
func main() {
	var s, sep string
	// os.Args是个切片
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("args:" + s)
}
