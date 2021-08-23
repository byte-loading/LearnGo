package main

import (
	"bufio"
	"fmt"
	"os"
)

// 参数中如果有值，读取文件中的数据，获取重复行，否则读取用户输入
func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("please input: ")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// 打开文件，并处理异常
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for fileName, v := range counts {
		for line, n := range v {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, n, line)
			}
		}
	}
}

// 统计文件中，重复行出现的次数
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "bye" {
			break
		}
		if counts[f.Name()] == nil {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
}
