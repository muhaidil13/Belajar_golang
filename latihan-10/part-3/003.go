package main

import "fmt"

func main() {
	// 6
	a := make(chan int)
	go func() {
		count := 100
		for i := 0; i < count; i++ {
			a <- i
		}
		close(a)
	}()
	for v := range a {
		fmt.Println(v)
	}
	// 7
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

}
