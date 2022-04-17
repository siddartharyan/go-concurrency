package first

import (
	"fmt"
	"time"
)

func Print(name string) {

	for i := 0; i < 3; i++ {
		fmt.Println(name)
		time.Sleep(1 * time.Millisecond)
	}
}
