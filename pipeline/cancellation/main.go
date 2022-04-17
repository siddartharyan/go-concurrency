package main

import (
	"fmt"
	"sync"
)

// generator() -> square() -> print()
func main() {

	done := make(chan struct{})
	ch := square(done, merge(done, generate(9, 8, 7), generate(1, 3, 4)))
	fmt.Println(<-ch)
	close(done)
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

func square(done <-chan struct{}, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for val := range ch2 {
			select {
			case ch <- (val * val):
			case <-done:
				return
			}

		}
	}()
	return ch
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func() {
			defer wg.Done()
			for x := range c {
				select {
				case ch <- x:
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
