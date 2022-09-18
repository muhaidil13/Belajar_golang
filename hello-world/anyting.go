package main

import (
	"fmt"
)

// Dieksekusi secara control over flow
func main() {
	fmt.Println("Test hello world")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			getName()
		}
	}
	n, _ := fmt.Println("nama_", 123, true)
	// jumlah keseluruhan
	fmt.Println(n)
	// errortest()
}

func getName() {
	fmt.Println("My Name Farid")
	nameUser()
}
func nameUser() {
	fmt.Println("-------------")
}
