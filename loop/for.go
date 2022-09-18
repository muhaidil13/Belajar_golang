package main

import (
	"fmt"
)

func main() {
	i := 10
	for i <= 20 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// infinite loop pakai break untuk berhenti
	// for {
	// 	fmt.Println("hello")
	// }

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("sudah sampai ", i)
			continue
		}
		fmt.Println(i)
	}
	// Loop Didalam Loop
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("Ini adalah scope luar %d\t ini adalah Scope dalam %d\n", i, j)
		}
	}
	// Membuat format acsi ke sting menggunakan for
	for i := 33; i < 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}

	// if Statement akan menghasilkan bolean

	number := 41
	if number == 10 && number%2 == 0 {
		fmt.Println("sama")
	} else if number != 10 && number%2 == 0 {
		fmt.Println("bukan 10")
	} else if number >= 20 || number < 20 {
		fmt.Println("number")
	} else {
		fmt.Println("null")
	}

	// switch akan memilih yang true
	// fallthrouh memungkinkan membuat next linenya di exekusi dengan syarat harus
	// true awalnya
	switch {
	case false:
		fmt.Println("its false case 1")
	case (3 == 2):
		fmt.Println("its true")
		fallthrough
	case (20 == 20):
		fmt.Println("20 tidak sama 20")
		fallthrough
	case (30 == 20):
		fmt.Println("30 tidak sama 20")
		fallthrough
	case (40 == 20):
		fmt.Println("40 tidak sama 20")
	case (50 == 20):
		fmt.Println("50 tidak sama 20")
	default:
		fmt.Println("this default")
	}
	n := "boby"
	switch n {
	case "mob":
		fmt.Println("this bobs")
	case "budi":
		fmt.Println("its budi")
	case "boby":
		fmt.Println("its boby")
	default:
		fmt.Println("its people")
	}

	fmt.Println(loops(10))

}
func loops(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}
