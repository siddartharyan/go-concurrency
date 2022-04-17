package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIdKeyType string

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	ctx1 := context.WithValue(ctx, userIdKeyType("userid"), "jane")
	ch := checkMemberShip(ctx1)

	fmt.Println("value for jane is ", <-ch)

}

func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)

	go func() {
		defer close(ch)

		userid := ctx.Value(userIdKeyType("userid")).(string)

		ch <- db[userid]
	}()

	return ch
}
