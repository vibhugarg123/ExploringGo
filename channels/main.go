package main

import (
	"fmt"
	"sync"
)

/*
	func main() {
		fmt.Println("Channels in golang")

		myCh := make(chan int)

		// It's always towards left
		myCh <- 5
		fmt.Println(<-myCh)
	}

	Channels are the way go routines talk to each other by passing miniature values.
	DO NOT COMMUNICATE BY SHARING MEMORY. INSTEAD, SHARE MEMORY BY COMMUNICATING.

	It gives error:= Channels in golang fatal error: all goroutines are asleep - deadlock!
	Channels only allow values passing to them when somebody is listening to them.

	Flipping line will work-
	fmt.Println(<-myCh)
	myCh <- 5

	* channels are bidirectional. You can send the value or read from channel.
	* SEND ONLY CHANNEL
		ch chan<- int
	* RECEIVE ONLY CHANNEL: Not allowed to put close() inside it
		ch <-chan int

	Buffered Channel: Specifies the capacity of channel
		ch:=make(chan int, 1)

*/

func main() {
	fmt.Println("Channels in golang")

	//Buffered Channel
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	//Receive ONLY from channel, <-chan going outside of the box
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// differentiate whether you are reading from a closed channel
		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println(val)
		}

		wg.Done()
	}(ch, wg)

	//Send ONLY to channel, chan<- going inside the box
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		// Pass value first and then close the channel, otherwise it panics: send on closed channel
		close(ch)

		wg.Done()
	}(ch, wg)

	wg.Wait()
}
