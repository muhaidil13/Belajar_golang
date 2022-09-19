package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Item  []string
}

func main() {
	// 1
	p1 := person{
		First: "Farid",
		Last:  "wajdi",
		Item:  []string{"1", "2", "4"},
	}
	c, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("Terdapat Error pada ", err)
	}
	fmt.Println(string(c))
}
