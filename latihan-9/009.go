package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 1
var wg sync.WaitGroup

// 2
type person struct {
	First string
	Last  string
}

type human interface {
	speak()
}

func (ps *person) speak() {
	fmt.Println("hello")
}
func saysomething(hm human) {
	hm.speak()
}
func main() {
	// // 1
	// fmt.Println("cpu awal", runtime.NumCPU())
	// fmt.Println("rountine awal", runtime.NumGoroutine())
	// fmt.Println("test b")
	// wg.Add(2)
	// go func() {
	// 	fmt.Println("This withe group 1")
	// 	wg.Done()
	// }()
	// go func() {
	// 	fmt.Println("This withe group 2")
	// 	wg.Done()
	// }()
	// fmt.Println("cpu mid", runtime.NumCPU())
	// fmt.Println("rountine mid", runtime.NumGoroutine())
	// wg.Wait()
	// fmt.Println("cpu akhir", runtime.NumCPU())
	// fmt.Println("rountine akhir", runtime.NumGoroutine())
	// fmt.Println("test a")

	// 2
	c := person{
		First: "farid",
		Last:  "wajdi",
	}
	// interface
	saysomething(&c)
	// method
	c.speak()

	// 3 waitgroub race condision
	// var mu sync.Mutex
	// ads := 0
	// asdd := 100
	// wg.Add(asdd)
	// for i := 0; i < asdd; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		v := ads
	// 		fmt.Println(ads)
	// 		runtime.Gosched()
	// 		v++
	// 		ads = v
	// 		fmt.Println(ads)
	// 		mu.Unlock()
	// 		wg.Done()
	// 	}()
	// }
	// // command go run -race main.go
	// wg.Wait()
	// fmt.Println("end", ads)

	// 4 mutex(skip) // atomic
	var at int64 = 0
	asdd := 100
	wg.Add(asdd)
	for i := 0; i < asdd; i++ {
		go func() {
			atomic.AddInt64(&at, 1)
			fmt.Println("nilai ", atomic.LoadInt64(&at))
			wg.Done()
			// fmt.Println(runtime.NumGoroutine())
		}()
	}
	// command go run -race main.go
	wg.Wait()
	fmt.Println("end", at)

	// 6 os
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	// command
	// 1 go build (untuk membuat apk sesuai sistem operasi liat pada folder yang sama)
	// jalankan ./nama file
	// maka hasilnya pasti sama dengan go run namafile.go
	// go install (membuatkan file apk yang disimpan pada folder $PATH)
}
