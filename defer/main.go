package main

import (
	"fmt"
	"time"
)

func main() {
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	for i := 0; i < 3; i++ {
		go fmt.Println("loopFunc:", i)
	}
}
