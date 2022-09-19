package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	recieve(even, odd, quit)
}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("ini dari even ", v)
		case v := <-o:
			fmt.Println("ini dari odd", v)
		case v := <-q:
			fmt.Println("ini dari quuit ", v)
			return
		}
	}
}
