package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Millisecond)
		ch <- 1
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(1 * time.Millisecond):
		fmt.Println("timeout")
	}
}
