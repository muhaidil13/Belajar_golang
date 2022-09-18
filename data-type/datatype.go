package main

// Bisa dibaca dokumentasinya
import "fmt"

var y = 20

var z = "Type String"

func main() {

	// Membuat type sendiri
	type nomor int

	var test nomor = 10
	fmt.Printf("%T\n", test) // Main.nomor type

	y = int(test)
	fmt.Printf("%T\t", test)

	fmt.Println(y)
	// type
	fmt.Printf("%T\n", y)
	// value
	fmt.Printf("%#v\n", y)

	// Binary
	fmt.Printf("%#b\n", y)

	// hexa
	fmt.Printf("%#x\n", y)

	fmt.Printf("----------------------------------------------------\n")

	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// Bisa Seperti ini
	fmt.Println(`nama "saya"`)

	// z sudah di assigment sebagai string Its static programing
	// z = 10
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
