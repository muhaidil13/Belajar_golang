package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Nilai default false
	var benar bool
	a := 10
	b := 10
	fmt.Println(a == b)
	fmt.Println(benar)

	// type int
	var nilai1 int16 //ini memiliki nilai
	nilai1 = 10
	fmt.Println(nilai1)

	// hanya memiliki nilai
	var nilai2 uint8
	nilai2 = 10
	fmt.Println(nilai2)

	fmt.Print(runtime.GOOS)
	fmt.Print(runtime.GOARCH)

	s := "test"
	fmt.Printf("%T\n", s)

	// Bilangan ASSCI
	sb := []uint8(s)
	sc := []byte(s)
	fmt.Println(sb)
	fmt.Println(sc)

	fmt.Printf("%T\n", sb)
	fmt.Printf("%T\n", sc)

	for i := 0; i < len(s); i++ {
		// type char
		fmt.Printf("%U", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("index %#U and hex %#x\n", i, v)
	}
}
