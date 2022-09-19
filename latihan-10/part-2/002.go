package main

import "fmt"

// lanjutan 4
func main() {

	// 4
	ab := make(chan int)
	c := gen(ab)
	receive(c, ab)

	// 5
	caa := make(chan int)
	go func() {
		caa <- 10
	}()
	v, ok := <-caa
	fmt.Println(v, ok)
	close(caa)
	v, ok = <-caa
	fmt.Println(v, ok)

}

// send
func gen(x chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		x <- 1
		close(c)
	}()
	return c
}

// receive
func receive(a, b <-chan int) {
	for {
		select {
		case v := <-a:
			fmt.Println(v)
		case <-b:
			return
		}
	}
}
