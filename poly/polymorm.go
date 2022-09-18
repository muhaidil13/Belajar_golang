package main

import "fmt"

// semua ini adalah polymorfism

type orang struct {
	nama string
	nim  int
}
type mahasiswa struct {
	orang
	ruangan string
	jurusan string
}
type mahasiswi struct {
	orang
	ruangan string
	jurusan string
	bawaan  string
}
type nilaidata int

// Membuat interface human berisi method konsul
type human interface {
	// method konsul apapun
	konsul()
}

func main() {
	org1 := mahasiswi{
		orang: orang{
			nama: "dila",
			nim:  10184007,
		},
		ruangan: "100",
		jurusan: "si",
		bawaan:  "dompet",
	}
	org2 := mahasiswa{
		orang: orang{
			nama: "farid",
			nim:  10184007,
		},
		ruangan: "102",
		jurusan: "SI",
	}
	// Menampung value dari objek
	bar(org1)
	bar(org2)

	// convertion
	var d nilaidata = 10
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	var y int
	y = int(d)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// membuat Fungsi Anonymous
	func() {
		fmt.Println("test anonymouse function")
		// kurung buka dibawah merupakan code exekusi dari function ini
	}()

	// Membuat fungsi dengan argumen
	func(nama string, nim int) {
		fmt.Println("nama saya adalah ", nama, "dan nim saya ", nim)
	}("farid", 10184007)

	// membuat function expression
	j := func() {
		fmt.Println("hallo bang")
	}

	// Membuat functions expression with arguments
	jq := func(x int) {
		fmt.Println("nilai yang saya masukkan adalah", x)
	}
	// cara panggil functions expressions dan yang ada argument
	j()
	jq(10)

	// string
	a := foo()
	fmt.Println(a)

	// functions
	// func namafunc kembalian(type data)

	as := test()
	fmt.Printf("%T\n", as)
	fmt.Println(as())
	fmt.Println(test()())

}

// mengembalikan fungsi
func test() func() int {
	return func() int {
		return 10000
	}
}

// Membuat method dengan keyword return
func foo() string {
	s := "hello world"
	return s
}

// membuat metod
func (s mahasiswi) konsul() {
	fmt.Println("nama saya ", s.nama, "ini adalah mahasiswi")
	fmt.Println("nim saya ", s.nim)
}
func (s mahasiswa) konsul() {
	fmt.Println("nama saya", s.nama, "ini adalah mahasiswa")
	fmt.Println("nim saya ", s.nim)
}

func bar(h human) {
	// memanggil semua data jika menggunakan method ini
	// menggunakan switch
	switch h.(type) {
	case mahasiswa:
		fmt.Println("ini fungsi human", h.(mahasiswa).nama)
	case mahasiswi:
		fmt.Println("ini fungsi human", h.(mahasiswi).nama)

	}
}
