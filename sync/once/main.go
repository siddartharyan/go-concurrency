package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Once is used to call function only once
	var wg sync.WaitGroup
	var once sync.Once

	load := func() {
		fmt.Println("should be called during initialization")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(load)
		}()
	}

	wg.Wait()
}
