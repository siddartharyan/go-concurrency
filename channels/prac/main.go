package main

import "fmt"

func main() {

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Println("received value is: ", v)
		}
		fmt.Println("Done with receiving")
	}

	ch := owner()
	consumer(ch)
}

func owner() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			ch <- i
		}
	}()
	return ch
}
