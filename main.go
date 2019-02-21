package main

import (
	"fmt"
)

func main() {
	var input string
	var flag bool = false
	for true {
		//	fmt.Printf("Json:")
		fmt.Scanf("%s", &input)
		flag = false
		// 输入exit或quit则退出
		if (input == "quit") || (input == "exit") {
			break
		}
		fmt.Printf("执行.")
	}
}
