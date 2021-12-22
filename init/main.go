package main

import (
	"fmt"

	_ "github.com/LittleMonsterZL/examples/init/a"
	_ "github.com/LittleMonsterZL/examples/init/b"
)

func init() {
	fmt.Println("main init")
}

func main() {

}
