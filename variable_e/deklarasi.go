package variable_e

import "fmt"

// diluar jangan gunakan short scope global
// var test = "coba"

// memiliki nilai default nilai 0 string kosong boolean false function dll null
// var test1 int

func variable_e() {
	// Deklarasi dengan identifier
	var b int = 10
	fmt.Print(b)

	// Variable local
	var a = "cotoh"
	fmt.Println(a)

	/*
		note jika variable sudah dideklarasi maka harus digunakan := hampir sama
		dengan let dan var pada js namum pada saat di deklarasi harus digunakan
		jika ingin mendeklarasikan saja akan dibahas berikutnya

	*/
	// Deklarasi variable dan assigin value (short)
	x := 42
	fmt.Println(x)

	// assiggin value ke variable new value
	x = 12
	fmt.Println(x)

	// Expressions adalah kombinasi explisit dengan value lain and statement
	// adalah perintah untuk dijalankan

	y := 12 + 10
	fmt.Println(y)
	fmt.Println(test)
	fmt.Println(test1)

}
