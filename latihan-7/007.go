package main

import "fmt"

type person struct {
	nama string
	nim  int
}

func main() {
	// 1
	a := 10
	fmt.Println(&a)
	fmt.Println(a)

	// 2
	p1 := person{
		nama: "farid",
		nim:  10184007,
	}
	fmt.Println(&p1.nama)
	fmt.Println(&p1.nim)
	change(&p1)
	fmt.Println(&p1.nim)
	fmt.Println(&p1.nama)
	fmt.Println(p1.nama)

}

func change(p *person) {
	(*p).nama = "farid wajdi"
	p.nim = 123232323

}
