package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	// 将函数定义为 map 的 value
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	// 访问 map
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	// 遍历 map
	for k, v := range myMap {
		println(k, v)
	}
}
