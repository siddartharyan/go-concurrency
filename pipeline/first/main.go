package main

import "fmt"

// generator() -> square() -> print()
func main() {
	ch := square(generate(1, 3, 4, 5, 6))
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
