package main

import "fmt"

func main() {
	// 1.1 go
	a := make(chan int)
	go func() {
		a <- 16
	}()
	fmt.Println(<-a)

	// 1.2 buffer
	b := make(chan int, 1)
	b <- 19
	fmt.Println(<-b)

	// 2 in out
	cs := make(chan int)
	go func() {
		cs <- 10
	}()
	fmt.Println(<-cs)

	// 3 assign value and pull with range loop
	as := lop()
	rev(as)

}

// send 3
func lop() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c

}

// receive 3
func rev(x <-chan int) {
	for v := range x {
		fmt.Println(v)
	}
}
