package main

import (
	"fmt"
)

func main() {
	ev := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(ev, odd, quit)
	receive(ev, odd, quit)

	// Test
	ssa := make(chan int)
	go func() {
		ssa <- 10
		close(ssa)
	}()
	v, ok := <-ssa
	fmt.Println(v, ok)
	// maka akan kosong jika sudah digunakan
	fmt.Println(<-ssa)

}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	q <- true
	close(q)
}
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("Data Even ", v)
		case v := <-o:
			fmt.Println("Data Odd ", v)
		case i, v := <-q:
			if !v {
				fmt.Println("test", i)
				return
			} else {
				fmt.Println("test else", i)
			}

		}
	}
}
