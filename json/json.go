package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
	Last  string
	Nim   int
}

// unMarshal
type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Nim   int    `json:"nim"`
}

// sort costume
type Sorting struct {
	Nama string
	Nim  int
	Age  int
}
type byname []Sorting

func (bn byname) Len() int {
	return len(bn)
}
func (bn byname) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}
func (bn byname) Less(i, j int) bool {
	// jika namanya upper case maka akan didahulukan setelah itu baru lowercase
	return bn[i].Nama < bn[j].Nama
}
func main() {
	p1 := person{
		First: "Farid Wajdi",
		Last:  "Anwar",
		Nim:   10184007,
	}
	p2 := person{
		First: "Alif",
		Last:  "Budi",
		Nim:   10184011,
	}
	pp := []person{p1, p2}
	fmt.Println(pp)
	// Json Marhal

	bs, err := json.Marshal(pp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// json Unmarshal
	s := `[{"First":"Farid Wajdi","Last":"Anwar","Nim":10184007},{"First":"Alif","Last":"Budi","Nim":10184011}]`
	sb := []byte(s)
	var people []person2
	// people := []person2{}
	err1 := json.Unmarshal(sb, &people)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Hasil ", people)
	for i, v := range people {
		fmt.Println("=============person", i)
		fmt.Println(v.First, v.Last, v.Nim)
	}

	// sort
	ab := []int{1, 2, 11, 22, 33, 10, 0, 10, 3, 4, 6}
	ac := []string{"fari", "wajdi", "human", "tes", "zebra", "ali", "neks"}
	// unsort
	fmt.Println(ab)
	sort.Ints(ab)
	// sort
	fmt.Println(ab)

	fmt.Println("==========================")
	// unsort
	fmt.Println(ac)
	sort.Strings(ac)
	// sort
	fmt.Println(ac)

	// ////////////////////////////////////////
	// Sorting costume
	s1 := Sorting{
		Nama: "Farid",
		Nim:  100000,
		Age:  10,
	}
	s2 := Sorting{
		Nama: "adi",
		Nim:  200000,
		Age:  12,
	}
	s3 := Sorting{
		Nama: "carid",
		Nim:  4400000,
		Age:  9,
	}
	s4 := Sorting{"sodi", 20, 10}
	test := []Sorting{s1, s2, s3, s4}
	fmt.Println(test)
	sort.Sort(byname(test))
	fmt.Println(test)

	// cari cara dlu
	// go get golang.org/x/crypto/bcrypt
	// bycrpit
	// bs, err := bcrypt.GenerateFormPassword([]byte(pass), bcrypt.MinCost)

	// hasing
	pass := `farid123`
	bs, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	pass1 := `farid1213`
	// compare
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass1))
	if err != nil {
		fmt.Println("you cant login")
		return
	}
	fmt.Println("You login")

}
