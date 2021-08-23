package main

import (
	"fmt"
	"os"
	"strings"
)

// 带执行路径的参数
func withPathArgs(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println("args: ", s)
}

// 不带执行路径，只有参数
func onlyArgs(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("args: ", s)
}

// args join
func argsJoin(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

// 遍历每一行
func echoLine(args []string) {
	fmt.Println("path: ", args[0])
	for i, v := range args[1:] {
		fmt.Println("【index:value】", "【", i, ":", v, "】")
	}
}

func main() {
	withPathArgs(os.Args)
	onlyArgs(os.Args)
	argsJoin(os.Args)
	echoLine(os.Args)
}
