package main

import (
	"fmt"
	"sync"
)

// fan in channel
func main() {
	ev := make(chan int)
	od := make(chan int)
	fanin := make(chan int)

	go send(ev, od)
	go receive(ev, od, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
}

// send channel
func send(ev, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ev <- i
		} else {
			odd <- i
		}
	}
	close(ev)
	close(odd)
}

// recive channel
func receive(ev, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range ev {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)

}
