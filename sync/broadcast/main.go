package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	go func() {
		defer wg.Done()
		cond.L.Lock()
		for len(sharedRsc) < 1 {
			cond.Wait()
		}

		fmt.Println("processing ... ", sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		cond.L.Lock()
		for len(sharedRsc) < 2 {
			cond.Wait()
		}

		fmt.Println("processing ... ", sharedRsc["rsc2"])
		cond.L.Unlock()
	}()

	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()
}
