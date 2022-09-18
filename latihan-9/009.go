package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 1
	fmt.Println("cpu awal", runtime.NumCPU())
	fmt.Println("rountine awal", runtime.NumGoroutine())
	fmt.Println("test b")
	wg.Add(2)
	go func() {
		fmt.Println("This withe group 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("This withe group 2")
		wg.Done()
	}()
	fmt.Println("cpu mid", runtime.NumCPU())
	fmt.Println("rountine mid", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("cpu akhir", runtime.NumCPU())
	fmt.Println("rountine akhir", runtime.NumGoroutine())
	fmt.Println("test a")
}
