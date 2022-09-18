package main

import (
	"fmt"
)

func main() {
	// Membuat array composite, range key value, print format
	// 1
	jq := [6]int{10, 20, 30, 40, 50}
	jq[5] = 60
	fmt.Println(jq)

	for k, v := range jq {
		fmt.Println(k, v)
	}
	fmt.Printf("%T\n", jq)

	// Slice
	// 2
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(j[:3])
	fmt.Println(j[3:])
	fmt.Println(j[3:6])

	// 3
	ji := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ji = append(ji, 11, 12, 13)
	fmt.Println(ji)

	j = append(j, ji...)
	fmt.Println(j)

	// 4 append and slice
	jj := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jj = append(jj[:3], jj[6:]...)
	fmt.Println(jj)

	// 5 make jumlah dan kapasitas max
	jw := make([]string, 10, 20)
	fmt.Println(jw)
	jw = []string{
		"satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat", "satu", "dua", "tiga", "empat",
	}
	fmt.Println(len(jw))
	fmt.Println(cap(jw))
	for i := 0; i < len(jw); i++ {
		fmt.Println(i, jw[i])

	}

	// 6 slice of slice
	d := []int{1, 2, 3, 4, 5, 6}
	w := []int{7, 8, 9}
	x := [][]int{d, w}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println("data ke ", i)
		for j, a := range v {
			fmt.Printf("index posision ke %v value%v\n", j, a)
		}
	}
	// map 7
	cd := map[string][]string{
		"nama":  []string{"farid", "alif"},
		"kelas": []string{"Si-RM43", "sirm43"},
		"nim":   []string{"10184007", "10121212"},
	}
	fmt.Println(cd["nama"])
	fmt.Println(cd)
	for i, v := range cd {
		fmt.Println("Key = ", i)
		for i, vd := range v {
			fmt.Printf("key %v dan value %v\n", i, vd)
		}
	}
	// 8 membuat map
	cd["jurusan"] = []string{"Sistem Informasi", "Management Informatika", "Akuntansi"}
	fmt.Println(cd)
	for i, v := range cd {
		fmt.Println("Key = ", i)
		for i, vd := range v {
			fmt.Printf("key %v dan value %v\n", i, vd)
		}
	}
	fmt.Println(cd)
	delete(cd, "jurusan")
	fmt.Println(cd)

}
