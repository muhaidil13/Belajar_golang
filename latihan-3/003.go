package main

import (
	"fmt"
)

func main() {
	// Membuat loop 1 sampai 10000
	// for i := 1; i <= 10000; i++ {
	// 	fmt.Println(i)
	// }

	// Membuat Loop dalam loop
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(i)
	// 	for j := 40; j <= 90; j++ {
	// 		fmt.Printf("\t%#U\n", j)
	// 	}
	// }

	// Membuat for dengan Kondisi
	// b := 10
	// for b <= 20 {
	// 	fmt.Println(b)
	// 	b++
	// }

	// Membuat loop for dengan gabungan if
	// nilai := 10
	// for {
	// 	if nilai >= 20 {
	// 		break
	// 	}
	// 	fmt.Println(nilai)
	// 	nilai++
	// }

	// membuat loop modulus
	// for i := 0; i <= 100; i++ {
	// 	fmt.Printf("when %v, modulus 4 menghasilkan  %v\n", i, i%4)
	// }

	// if eslse statemenet
	// a := "farid"
	// if a == "aidil" {
	// 	fmt.Println("aidil ini orangnya")
	// } else if a == "budi" {
	// 	fmt.Println("budi ini orangnya")
	// } else {
	// 	fmt.Println("ini orangnya")
	// }

	// membuat switch
	switch {
	case false:
		fmt.Println("f")
	case true:
		fmt.Println("t")
	}
	// membuat switch
	hoby := "bola"
	switch hoby {
	case "ninja":
		fmt.Println("hobi saya ninja")
	case "lari":
		fmt.Println("hobi saya lari")
	case "bola":
		fmt.Println("hobi saya bola")
	default:
		fmt.Println("saya tidak punya hobi")
	}

}
