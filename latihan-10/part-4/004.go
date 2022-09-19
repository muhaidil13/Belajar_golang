package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	a := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-a)
		fmt.Println("Gorountine = ", runtime.NumGoroutine())

	}
}

func gen(x, y int) <-chan int {
	c := make(chan int)
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("Gorountine = ", runtime.NumGoroutine())
	}
	return c
}
