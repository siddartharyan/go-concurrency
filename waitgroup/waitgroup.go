package main

import (
	"fmt"
	"sync"
)

func main() {

	var data int
	var wg sync.WaitGroup
	wg.Add(1) //no of goroutines we are creating

	go func() {
		defer wg.Done() //mark it as done once our work is done
		data++
	}()

	wg.Wait() //waits until it is marked as done

	fmt.Println("The value of data is", data)
	fmt.Println("done...")
}
