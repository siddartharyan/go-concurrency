package main

import (
	"context"
	"fmt"
)

func main() {

	generator := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		i := 1
		go func() {
			defer close(ch)
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
					i++
				}
			}
		}()
		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := generator(ctx)

	for i := range ch {
		fmt.Println(i)

		if i == 5 {
			cancel()
		}
	}

}
