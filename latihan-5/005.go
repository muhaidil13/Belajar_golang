package main

import "fmt"

// 1
type person struct {
	nama  string
	kelas string
	skill []string
}

func main() {
	p1 := person{
		nama:  "farid",
		kelas: "sirm54",
		skill: []string{"html", "css", "golang", "javascrpit", "dll"},
	}
	p2 := person{
		nama:  "alif",
		kelas: "sirm24",
		skill: []string{"editing", "css", "golang", "javascrpit", "dll"},
	}
	fmt.Println(p1)
	for k, v := range p1.skill {
		fmt.Println(k, v)
	}

	// 2

	m := map[string]person{
		p1.nama: p1,
		p2.nama: p2,
	}
	fmt.Println(m)
	for i, k := range m {
		fmt.Println("id of ", i, k)
		for i, v := range p1.skill {
			fmt.Println(i, v)
			fmt.Println("=========================")
		}
	}

	// 3
	type kendaraan struct {
		doors int
		color []string
	}

	type motor struct {
		kendaraan
		roda bool
	}

	type mobil struct {
		kendaraan
		nitro bool
	}
	sedan := mobil{
		kendaraan: kendaraan{
			doors: 4,
			color: []string{"merah", "biru"},
		}, nitro: true,
	}
	fmt.Println(sedan)

	honda := motor{
		kendaraan: kendaraan{
			doors: 1,
			color: []string{"red", "black"},
		},
		roda: true,
	}
	fmt.Println(honda)
	fmt.Println(honda.doors)

	s := struct {
		nama  string
		umur  int
		skill []string
	}{
		nama:  "farid",
		umur:  21,
		skill: []string{"html", "css", "golang"},
	}
	fmt.Println(s.nama)
	fmt.Println(s.umur)
	for k, v := range s.skill {
		fmt.Println(k, v)
	}

	users := struct {
		users map[string]string
		level []string
	}{
		users: map[string]string{
			"admin":  "sebagai admin",
			"gudang": "digudang",
		},
		level: []string{
			"1", "2", "3",
		},
	}
	fmt.Println(users)
	for i, v := range users.users {
		fmt.Println(i, v)
	}
	for i, v := range users.level {
		fmt.Println(i, v)
	}
}
