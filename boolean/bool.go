package main

import "fmt"

// const (
// 	nama  string = "farid"
// 	kelas string = "SI-RM53"
// )
const (
	// Incement By 1
	a = iota
	b
	c
	d
)

func main() {
	data := fmt.Sprintln("data")
	fmt.Println(data)
	as := 10
	ass := 20
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(a == b)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(as == ass)
}
