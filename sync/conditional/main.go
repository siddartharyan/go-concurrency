package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)

	go func() {
		defer wg.Done()
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait() // wait will release the lock and suspend go routine
		}
		fmt.Println("processing ...", sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	cond.L.Lock()
	sharedRsc["rsc1"] = "bar"
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()

}
