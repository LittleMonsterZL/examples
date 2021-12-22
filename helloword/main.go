package main

import (
	"fmt"
)

// main 函数
func main() {
	err, result := DuplicateString("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

// 普通函数
func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}
