package main

import (
	"fmt"
	"sync"
)

func main() {

	waitGroup1()
	waitGroup2()
}

func waitGroup1() {
	var wg sync.WaitGroup

	fnc := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Println("the value is :", i)
		}()
		fmt.Println("returned from function")
	}
	fnc(&wg)
	wg.Wait()
	fmt.Println("function ended")
}

func waitGroup2() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
