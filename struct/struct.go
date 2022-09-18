package main

import (
	"fmt"
)

// Struct hampir sama dengan map tetapi dapat menammpung data
// dengan data tipe yang berbeda

type nama struct {
	depan    string
	balakang string
	age      int
}
type mahasiswa struct {
	// nama merupakan type dari user bertype stuct
	user  nama
	nilai int
}

func main() {

	p1 := nama{
		depan:    "farid",
		balakang: "wajdi",
		age:      21,
	}
	p2 := nama{
		depan:    "alif",
		balakang: "budi",
		age:      20,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.balakang)

	// Kita juga dapat emmbet struct
	// Cara menggunakannya seperti berikut
	p3 := mahasiswa{
		// user adalah field yang didefinisikan sebelumnya
		user: nama{
			depan:    "farid",
			balakang: "wajdi",
			age:      22,
		},
		nilai: 10,
	}
	fmt.Println(p3)
	fmt.Println(p3.user)

	// composite literal maksudnya deklarasi type langsung
	// pada saat deklarasi / anonymous struct tidak punya nama
	// type
	p22 := struct {
		agent string
		no    int
	}{
		// value
		agent: "Farid",
		no:    1,
	}
	fmt.Println(p22)

}
