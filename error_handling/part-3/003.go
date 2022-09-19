package main

import (
	"fmt"
)

func main() {
	fu()
	fmt.Println("Normal-----")
}
func fu() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover di fu", r)
		}
	}()
	fmt.Println("Calling Next")
	next(0)
	fmt.Println("Selesai memanggil Next")
}
func next(i int) {
	if i > 3 {
		fmt.Println("Panic ------")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in next", i)
	fmt.Println("Print in next", i)
	next(i + 1)
}
