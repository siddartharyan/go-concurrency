package main

import "fmt"

/*
 <- chan string receive only channel
 chan <- string send only channel
goroutine which creates and writes the channel is the
owner of the channel and is responsible for closing the channel
*/
func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	//spin go routines

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	msg := <-ch2
	fmt.Println(msg)

}

func genMsg(ch1 chan<- string) {
	//send msg on ch1
	ch1 <- "hello world"

}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	//receive msg on ch1
	msg, _ := <-ch1
	//send it on ch2
	ch2 <- msg
}
