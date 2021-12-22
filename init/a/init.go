package a

import (
	"fmt"

	_ "github.com/LittleMonsterZL/examples/init/b"
)

func init() {
	fmt.Println("init from a")
}
