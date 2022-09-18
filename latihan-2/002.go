package main

import "fmt"

func main() {
	a := 51
	// Membuat format desimal binary hex
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)

	n := 10
	m := 21
	// Operasi Perbandingan
	b := (n == m)
	c := (n <= m)
	d := (n >= m)
	e := (n != m)
	f := (n > m)
	g := (n < m)
	fmt.Println(b, c, d, e, f, g)

	// const Bisa tidak Digunakan beda dengan var
	const (
		num      = 10
		num1 int = 20
	)
	fmt.Println(num, num1)

	// Membuat bilangan format desimal binary hex kemudian di shift
	number := 32
	fmt.Printf("%d\t%b\t%#x\t", number, number, number)
	fmt.Println("")
	// digeser 1 bit
	numberv := number >> 1
	fmt.Printf("%d\t%b\t%#x\t", numberv, numberv, numberv)

	// String raw literal
	text1 := `ini 
	adalah 
	string raw literas`
	fmt.Println(text1)

	// membuat incement tahun menggunakan iota dan constans
	const (
		t1 = 2000 + iota
		t2 = 2000 + iota
		t3 = 2000 + iota
		t4 = 2000 + iota
	)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

}
