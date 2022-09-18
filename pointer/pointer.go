package main

import "fmt"

func main() {
	// pointer / alamat penyimpanan pada memory
	var a int = 20
	fmt.Printf("%T\n", &a) // int
	fmt.Println(&a)        //sama dengan b

	// untuk mengambil nilai alamat suatu variable ditandai dengan &
	// untuk assigmen ditandai dengan * pada type data
	// & itu adalah poiner value
	// * itu adalh pointer type

	var b *int = &a
	fmt.Println(b)

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(&a)
	fmt.Println(&b)

	fmt.Println("===============================")
	x := 100
	fmt.Println("before", x)  //isi
	fmt.Println("before", &x) //tempat penyimpanan
	foo(&x)
	fmt.Println("after", x)
	fmt.Println("after", &x)

	abb := 10
	fmt.Println(abb)
	fmt.Println(&abb)
	poin(&abb)
	fmt.Println(&abb)
}

func foo(y *int) {
	fmt.Println("before", y)
	fmt.Println("before", *y)
	*y = 1000
	fmt.Println("after", y)
	fmt.Println("after", *y)
}

func poin(y *int) {
	fmt.Println("ini nilai", *y)
	*y = 100
	fmt.Println("ini nilai", y)
}
