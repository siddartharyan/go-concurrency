package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Millisecond)
			ch <- "message sent"
		}
	}()

	for i := 0; i < 2; i++ {

		select {
		case c := <-ch:
			fmt.Println(c)
		default:
			fmt.Println("no message")
		}
		fmt.Println("done with processing")
		time.Sleep(2 * time.Millisecond)
	}
}
