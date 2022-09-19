package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1 arah
func main() {
	c := fanin(boring("farid"), boring("wajdi"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("selesai")

}

// reciver
func boring(nm string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", nm, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}

// receiver
func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}

	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
