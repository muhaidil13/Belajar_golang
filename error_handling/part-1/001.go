package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Pengecekan error
func main() {
	// 1
	n, err := fmt.Print("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// 2
	var j1, j2, j3 string
	fmt.Print("name:")
	_, err = fmt.Scan(&j1)
	if err != nil {
		panic(err)
	}
	fmt.Print("kelas:")
	fmt.Scan(&j2)
	if err != nil {
		panic(err)
	}
	fmt.Print("nim:")
	fmt.Scan(&j3)
	if err != nil {
		panic(err)
	}
	fmt.Println(j1, j2, j3)

	// 3
	f, err := os.Create("Nama_Siswa")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	a := strings.NewReader("test")
	io.Copy(f, a)

	// 4
	fi, err := os.Open("Nama_Siswa")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fi.Close()
	b, err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
