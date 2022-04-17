package main

import (
	"fmt"
	"sync"
)

// generator() -> square() -> print()
func main() {
	ch := square(merge(generate(9, 8, 7), generate(1, 3, 4)))
	for val := range ch {
		fmt.Println("value is :", val)
	}
}

func generate(nums ...int) <-chan int {

	ch1 := make(chan int)
	go func() {
		for _, x := range nums {
			ch1 <- x
		}
		defer close(ch1)
	}()
	return ch1
}

func square(ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for val := range ch2 {
			ch <- (val * val)
		}
		defer close(ch)
	}()
	return ch
}

func merge(cs ...<-chan int) <-chan int {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func() {
			for x := range c {
				ch <- x
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
