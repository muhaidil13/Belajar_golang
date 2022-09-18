package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("os\n", runtime.GOOS)
	fmt.Println("arch\n", runtime.GOARCH)
	fmt.Println("cpus\n", runtime.NumCPU())
	fmt.Println("gorountime\n", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("cpus\n", runtime.NumCPU())
	fmt.Println("gorountime\n", runtime.NumGoroutine())
	wg.Wait()

	// chanel
	ch := make(chan int)
	go func() {
		ch <- tes(10)
	}()
	fmt.Println(<-ch)

	fmt.Println("==================================")
	// race condision
	// membuat kondisi pertukaran item dan bertambah 1
	// fmt.Println("CPUS = ", runtime.NumCPU())
	// fmt.Println("GOrountime =", runtime.NumGoroutine())
	// counter := 0
	// var gwr sync.WaitGroup
	// const ww = 100
	// gwr.Add(ww)
	// for i := 0; i < ww; i++ {
	// 	go func() {
	// 		v := counter
	// 		runtime.Gosched()
	// 		v++
	// 		counter = v
	// 		gwr.Done()
	// 	}()
	// 	fmt.Println("Goruountime ", runtime.NumGoroutine())
	// }
	// gwr.Wait()
	// fmt.Println("Goruountime ", runtime.NumGoroutine())
	// fmt.Println("count", counter)

	// mutex 100%
	// fmt.Println("CPUS = ", runtime.NumCPU())
	// fmt.Println("GOrountime =", runtime.NumGoroutine())
	// counter := 0
	// var gwr sync.WaitGroup
	// var mu sync.Mutex
	// const ww = 100
	// gwr.Add(ww)
	// for i := 0; i < ww; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		v := counter
	// 		runtime.Gosched()
	// 		v++
	// 		counter = v
	// 		mu.Unlock()
	// 		gwr.Done()
	// 	}()
	// 	fmt.Println("Goruountime ", runtime.NumGoroutine())
	// }
	// gwr.Wait()
	// fmt.Println("Goruountime ", runtime.NumGoroutine())
	// fmt.Println("count", counter)

	fmt.Println("CPUS = ", runtime.NumCPU())
	fmt.Println("GOrountime =", runtime.NumGoroutine())
	counter := 0
	var dat int64
	var gwr sync.WaitGroup
	const ww = 100
	gwr.Add(ww)
	for i := 0; i < ww; i++ {
		go func() {
			atomic.AddInt64(&dat, 1)
			runtime.Gosched()
			fmt.Println("counter", atomic.LoadInt64(&dat))
			gwr.Done()
		}()
		fmt.Println("Goruountime ", runtime.NumGoroutine())
	}
	gwr.Wait()
	fmt.Println("Goruountime ", runtime.NumGoroutine())
	fmt.Println("count", counter)
}

// courency dan paralism cari perbedaannya
// courency dapat memungkinkan code di excekusi secara bersamaan
func foo() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("fooo", i)
	}
	wg.Done()
}
func bar() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("bar", i)
	}
}
func tes(nu int) int { return nu * 3 }
