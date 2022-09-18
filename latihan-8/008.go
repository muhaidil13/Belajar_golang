package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type person struct {
	First string
	Last  string
}
type secret struct {
	First string
	Item  []string
	Last  string
	Age   int
}
type encode1 struct {
	First string   `json:First`
	Item  []string `json:Item`
	Last  string   `json:Last`
	Age   int      `json:Age`
}

type no5 []secret

func (x no5) Len() int {
	return len(x)
}
func (x no5) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
func (x no5) Less(i, j int) bool {
	return x[i].Age < x[j].Age
}

type namefirst []secret

func (xp namefirst) Len() int           { return len(xp) }
func (xp namefirst) Swap(i, j int)      { xp[i], xp[j] = xp[j], xp[i] }
func (xp namefirst) Less(i, j int) bool { return xp[i].First < xp[j].First }

func main() {
	// 1
	p1 := person{
		First: "farid",
		Last:  "wajdi",
	}
	p2 := person{
		First: "aidissl",
		Last:  "anwar",
	}
	p3 := person{
		First: "Rika",
		Last:  "Anwar",
	}
	users := []person{p1, p2, p3}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// 2
	s1 := secret{
		First: "aidsil1",
		Last:  "anwar",
		Item:  []string{"kaca", "pulpen", "tinta"},
		Age:   16,
	}
	s2 := secret{
		First: "farid",
		Last:  "wajdi",
		Item:  []string{"kaca", "pulpen", "tinta"},
		Age:   212,
	}
	s3 := secret{
		First: "rika",
		Last:  "wajdi",
		Item:  []string{"kaca", "pulpen", "tinta"},
		Age:   17,
	}
	s4 := secret{
		First: "aidsil",
		Last:  "anwar",
		Item:  []string{"kaca", "pulpen", "tinta"},
		Age:   126,
	}
	orangrumah := []secret{s1, s2, s3, s4}
	bs, err = json.Marshal(orangrumah)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	dara := `[{"First":"aidsil1","Item":["kaca","pulpen","tinta"],"Last":"anwar","Age":16},{"First":"farid","Item":["kaca","pulpen","tinta"],"Last":"wajdi","Age":22},{"First":"rika","Item":["kaca","pulpen","tinta"],"Last":"wajdi","Age":16},{"First":"aidsil","Item":["kaca","pulpen","tinta"],"Last":"anwar","Age":16}]`
	sj := []byte(dara)
	var orangan []encode1
	fmt.Println(dara)
	err = json.Unmarshal(sj, &orangan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orangan)
	for _, v := range orangan {
		fmt.Println(v.First, v.Last)
		for _, v := range v.Item {
			fmt.Println("\t", v)
		}
	}

	// 3 new encode and encode json
	json.NewEncoder(os.Stdout).Encode(orangrumah)

	// 4 sort
	ad := []int{1, 2, 11, 10, 20, 4, 5, 8, 200, 9, 4}
	sort.Ints(ad)
	// berdasarkan int
	fmt.Println(ad)

	// berdasarkan string
	ac := []string{"z", "b", "f", "c", "e", "h", "j"}
	sort.Strings(ac)
	fmt.Println(ac)

	// 5 sort by age
	nomor5 := []secret{s1, s2, s3, s4}
	fmt.Println(nomor5)
	sort.Sort(no5(nomor5))
	for _, va := range nomor5 {
		fmt.Println(va.First, va.Last, va.Age)
		for _, item := range va.Item {
			fmt.Println("\t", item)
		}
	}

	// sort by namefirst
	sort.Sort(namefirst(nomor5))
	fmt.Println(nomor5)
	for _, v := range nomor5 {
		fmt.Println(v.First)
		for _, v := range v.Item {
			fmt.Println("\t", v)
		}
	}

}
