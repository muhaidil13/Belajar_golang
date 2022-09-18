package main

import "fmt"

var aa string
var ab int
var ac bool

func main() {
	a := "String"
	b := 12
	c := true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a, b, c)
	fmt.Println("=======type============")
	fmt.Printf("%T\n", aa)
	fmt.Printf("%T\n", ab)
	fmt.Printf("%T\n", ac)

	fmt.Println("========value============")
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	// format banyak
	ad := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(ad)

	fmt.Println("========type define============")
	type tes int
	var key tes
	fmt.Println(key)
	fmt.Printf("%T\n", key)
	key = 10
	fmt.Printf("%v\n", key)
	type coba string
	var data coba
	fmt.Println(data)
	fmt.Printf("%T\n", data)
	data = "nama saya farid"
	fmt.Println(data)
	fmt.Println("========type convert============")
	data1 := data
	fmt.Println(data1)
	fmt.Printf("%T\n", data1)

}
