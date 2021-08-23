package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打印输入行包含的重复行
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("please input: ")
	for input.Scan() {
		if input.Text() == "bye" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
