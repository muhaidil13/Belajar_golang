package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ctx := context.Background()
	// fmt.Println("contex \t", ctx)
	// fmt.Println("contex err \t", ctx.Err())
	// fmt.Printf("context type : %T\t\n", ctx)

	// fmt.Println("++++++++++++++++++++++++")
	// ctx, cancel := context.WithCancel(ctx)

	// fmt.Println("contex \t\n", ctx)
	// fmt.Println("contex err \t", ctx.Err())
	// fmt.Printf("context type : %T\t", ctx)
	// fmt.Println("Cancel \t", cancel)
	// fmt.Printf("cancel type : %T\t", cancel)
	// fmt.Println("++++++++++++++++++++++++++++++")
	// cancel()
	// fmt.Println("contex \t\n", ctx)
	// fmt.Println("contex err \t", ctx.Err()) //setelah method cancel berubah jadi context cancel
	// fmt.Printf("context type : %T\t", ctx)
	// fmt.Println("Cancel \t", cancel)
	// fmt.Printf("cancel type : %T\t", cancel)
	// test doang
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("cek error ", ctx.Err())
	fmt.Println("Cek Rountine ", runtime.NumGoroutine())
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)

			}

		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Cek Error ke 2 ", ctx.Err())
	fmt.Println("Cek Rountine ke 2 ", runtime.NumGoroutine())
	fmt.Println("Akan di cancel")
	cancel()
	fmt.Println("di cancel")
	time.Sleep(time.Second * 2)
	fmt.Println("Cek Error ke 3 ", ctx.Err())
	fmt.Println("Cek Rountine ke 3 ", runtime.NumGoroutine())

}
