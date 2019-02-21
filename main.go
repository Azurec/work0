package main

import (
	"fmt"
)

func main() {
	for true {
		fmt.Printf("Json:")
		var input string
		fmt.Scanf("%s", &input)
		// 输入exit或quit则退出
		if (input == "quit") || (input == "exit") {
			break
		}
		fmt.Printf("执行.")
	}
}
