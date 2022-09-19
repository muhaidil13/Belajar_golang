package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go populate(c1)
	go fanout(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
func fanout(q <-chan int, e chan<- int) {
	var wq sync.WaitGroup
	for v := range q {
		wq.Add(1)
		go func(v2 int) {
			e <- waktu(v2)
			wq.Done()
		}(v)
	}
	wq.Wait()
	close(e)
}

func waktu(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
