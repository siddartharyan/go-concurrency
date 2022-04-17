package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "first routine"
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "second routine"
			time.Sleep(2 * time.Millisecond)
		}
	}()

	for {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
		}
	}

}
