package main

import (
	"fmt"
	"github.com/concurrency/first"
	"time"
)

func main() {
	//go routines
	routine()

	//channels

}

func routine() {
	first.Print("direct call")

	//Write go routine with different variations of function call

	// go routine function call

	go first.Print("goroutine-1")

	// go routine with anonymous function

	go func() {
		first.Print("anonymous goroutine")
	}()

	// go routine with function values call

	fv := first.Print

	go fv("value goroutine")

	// add sleep for go routines
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done ...")
}
