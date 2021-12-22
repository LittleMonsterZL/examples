package main

import "fmt"

func main() {
	a := []int{}
	b := []int{1, 2, 3}
	c := a
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	a = append(b, 1)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
