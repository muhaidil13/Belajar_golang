package main

import (
	"fmt"
)

type orang struct {
	nama string
	nim  string
}
type mahasiswa struct {
	orang
	nilai int
}

func main() {
	org1 := mahasiswa{
		orang: orang{
			nama: "farid",
			nim:  "1010101010",
		},
		nilai: 100,
	}
	org2 := mahasiswa{
		orang: orang{
			nama: "alif",
			nim:  "1sss101010",
		},
		nilai: 100,
	}
	fmt.Println(org1)
	fmt.Println(org2)

	// Akan dieksekusi pada akhir code
	defer foo()

	bar("farid")
	fmt.Println(tambah(12, 10))
	x, y := fmt.Println("betul", "false")
	fmt.Println(x)
	fmt.Println(y)
	// trueordare("true", "dare")
	a := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("hasil ", a)

	data := []int{1, 2, 3, 4, 6, 7, 8, 9}
	// digunakan pada array slice
	kal := kali(data...)

	fmt.Println(kal)

	// Membuat method harus membuat objek terlebih dahulu /
	// type struct dan harus di sisi data
	// pada saat deklarasi method harus disertai dengan objek di sela func
	// dan dapat ditambahkan keyword retrun dengan menentukan type data
	org1.bicara()

}

// method dengan objek disela func
// jika ingin menambah kembalian / keywird return juga bisa
func (s mahasiswa) bicara() {
	fmt.Println("nama saya adalah ", s.nama)
	fmt.Println("nilai saya adalah ", s.nilai)

}

// percobaan defer
func foo() {
	fmt.Println("Exit file")
}

// defer digunakan untuk mengakhiri sebuah code dan akan dieksekusi pada akhir
// dan dapat di letakkan dimana saja

// jika kosong maka default 0
// parameter array slice
func kali(x ...int) int {
	d := 0
	for i, v := range x {
		d = i
		d = d * v
		fmt.Println("index ", i, "dengan value ", v, "dengan value", d)
	}
	return d
}

// jika membuat keyword return maka harus di set dlu type valuenya

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("menambah value index", i, " dengan value ", v, "dengan ", sum)
	}
	fmt.Println("hasil dari Penjumlahan = ", sum)
	return sum
}

// format parameter
func bar(s string) {
	fmt.Println("Hello", s)
}

// Mendefinisikan kembalian
func tambah(a int, b int) int {
	return a + b
}

// membuat return berbeda
// func trueordare(t string, d string) (string, bool) {
// 	a := fmt.Sprintln(t, d, `"Berkata jangan"`)
// 	b := false
// 	return a, b
// }
