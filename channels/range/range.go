package main

import "fmt"

func main() {
	unbufferered()
	buffered()
}

func buffered() {
	ch := make(chan int, 5)

	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func unbufferered() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		defer close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
