package main

import (
	"fmt"
)

func main() {

	// Array pada golang kita dapat dengan mudah menentukan
	// panjang dan type dari sebuah array contoh
	var x [4]int
	// jika di print maka disi dari 4 index array di atas adalah 0
	fmt.Println(x)

	// kita juga dapat dengan mudah assign data pada array dengan index tertentu
	x[1] = 10
	fmt.Println(x)
	// panjang data
	fmt.Println(len(x))

	// grouping data
	// variable := [] type{values} //composite literals
	y := []int{1, 2, 3, 410, 10}
	fmt.Println(y)

	// looping array (i itu index v value)
	for i, v := range y {
		fmt.Println(i, v)
	}

	// Slice data pada array
	c := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(c)
	// slice dari sampai index
	fmt.Println(c[:])
	// slice dari 2 sampai semua
	fmt.Println(c[2:])
	// slice dari 2 sampai 5
	fmt.Println(c[2:5])

	// Append
	// kita coba tambah data c
	c = append(c, 10, 22, 11, 1022)
	d := []int{1, 2, 3, 4, 11, 2, 121, 121, 121323}
	fmt.Println(c)
	// Menggabungkan semua data array d
	c = append(c, d...)
	fmt.Println(c)

	// Memilih isi array slice delete form array
	c = append(c[:2], c[10:]...)
	// Maksud dari dibawah adalah dengan memiilh semua sampai index ke2
	// dan mengambil index ke10 sampai selesai
	fmt.Println(c)

	// Kita juga dapat membuat array menggunakan perintah make
	testarray := make([]int, 10, 1000)
	fmt.Println(testarray)
	// panjang array
	fmt.Println(len(testarray))
	// kapasitas maksimal array
	fmt.Println(cap(testarray))

	// Multipe dimensional array
	jb := []string{"Satu", "Dua", "Tiga"}
	jc := []string{"Empat", "Lima", "Enam"}
	xd := [][]string{jb, jc}
	fmt.Println(xd)

	// Map Array Multidimensional key and value
	jq := map[string]string{
		"satu": "test",
		"dua":  "test2",
	}
	fmt.Println(jq)
	// ini akan menghasilkan value dari keynya ingat jika value tidak ada maka
	// akan mengembalikan nilai defaultnya
	// fmt.Println(jq["satu"])
	key, val := jq["satu"]
	fmt.Println(key)
	// Value ini menghasilkan nilai true
	fmt.Println(val)

	// menggunakan if statement dengan array assoc
	if k, v := jq["sat"]; v {
		// akan mengembalikan kosong karena defaultnya kosong
		fmt.Println("nilai anda", k)
	}
	jq1 := map[string]int{
		"satu": 12,
		"dua":  13,
	}
	if ki, vv := jq1["satu"]; vv {
		// akan mengembalikan kosong karena defaultnya kosong
		fmt.Println("nilai anda", ki)
	}

	// Cara Menambahkan Element array k, v dan range loop
	jq1["tiga"] = 14
	fmt.Println(jq1)
	for f, v := range jq1 {
		fmt.Println(f, v)
	}

	// mendelete Array key value
	delete(jq1, "tiga")
	fmt.Println(jq1)
	if v, ok := jq1["empat"]; ok {
		fmt.Println("value:", v)
		delete(jq1, "empat")
	}
	fmt.Println(jq1)

}
