package main

import "fmt"

func main() {
	// call back adalah function sebagai argumen
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// s := sum(n...)
	// fmt.Println(s)

	// berdasarkan function contok even dia mengambil parameter di n
	// kemudian sum dijalankan menggunakan parameter yang dikembalikan oleh
	// even fungsi tadi
	// sum menjumlahkan
	t := even(sum, n...)
	fmt.Println(t)
	{
		// Skope
		y := 10
		fmt.Println(y)
	}
	// skope
	// fmt.Println(y_)
	od := odd(sum, n...)
	fmt.Println(od)

	// closure / fungsi tanpa nama
	a := incr()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// Recursion merupakan fungsi yang memanggil dirinya sendiri
	// harus menggunakan condisi supaya bisa selesai

	fmt.Println(perkalian(4))
	fmt.Println(loops(10))
}
func loops(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}

// contoh recursion
func perkalian(n int) int {
	if n == 1 {
		return 1
	}
	return n * perkalian(n-1)
}

// Mengembalikan fungsi
func incr() func() int {
	var x int
	// closure
	return func() int {
		x++
		return x
	}
}

// function with return value
func sum(num ...int) int {
	nilai := 0
	for _, v := range num {
		nilai += v
	}
	return nilai
}

// f func(xi ...int) callback di assigin ke f
//  xx ...int) int merupakan parameter
func even(f func(xi ...int) int, xx ...int) int {
	var yi []int
	for _, v := range xx {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
func odd(f func(xi ...int) int, xx ...int) int {
	var yi []int
	for _, v := range xx {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
