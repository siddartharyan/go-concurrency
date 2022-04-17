package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Millisecond))
	defer cancel()
	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)

		go func() {
			defer close(ch)

			deadline, ok := ctx.Deadline()

			if ok {
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("deadline exceeded")
					return
				}
			}
			select {
			case ch <- data{"hello"}:
			case <-ctx.Done():
				return

			}
		}()

		return ch
	}

	d, ok := <-compute(ctx)
	if ok {
		fmt.Println("received from channel ", d.result)
	} else {
		fmt.Println("exceeded")
	}

}
