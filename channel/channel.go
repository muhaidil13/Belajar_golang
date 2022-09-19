package main

import "fmt"

func main() {
	// buffer channel (menyediakan 1 ruang)
	c := make(chan int, 1)
	c <- 40
	fmt.Println(<-c)

	// buffer channel (menyediakan 2 ruang)
	e := make(chan int, 2)
	e <- 40
	e <- 41
	fmt.Println(<-e)
	fmt.Println(<-e)

	// Channel Gorountime block
	d := make(chan int)
	go func() {
		d <- 10
	}()
	fmt.Println(<-d)

	// directonal channel
	// channel dapat melakukan 1 aksi jika menerima maka hanya bisa menerima dan sebaliknya
	// receive dan send
	// dd := make(chan int)
	// rc := make(<-chan int) //reveive
	// sd := make(chan<- int) //send

	// tidak Bisa
	// note (spesifik ke general ) tidak bisa
	// dd = rc

	// spesifik ke spesifik tidak bisa
	// rc = sd

	// type Spesifik ke general tidak bisa
	// fmt.Printf("%T", (chan int)(rc))

	// bisa
	// general ke spesifik bisa
	// rc = dd
	// fmt.Printf("%T", rc) // <-chan int

	// type general spesifik conversi
	// fmt.Printf("%T", (<-chan int)(dd)) // <-chan int

	// Penggunaan Channel
	// dd := make(chan int)
	// go bar(dd)
	// foo(dd)
	// fmt.Println("Selesai")

	// contoh lain misalkan menggunakan loop
	test1 := make(chan int)
	// menyiapkan value
	go func() {
		for i := 0; i <= 100; i++ {
			test1 <- i
		}
		// harus di tutup jika tidak maka deadlock
		close(test1)
	}()
	// mengambil value
	for v := range test1 {
		fmt.Println(v)
	}

	// select

}

// // send hanya bisa push value
// func bar(c chan<- int) {
// 	c <- 100
// close(c)
// }

// // receiver (bisa baca, menerima value)
// func foo(c <-chan int) {
// 	fmt.Println(<-c)
// }
