package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Item  []string
}

func main() {
	p1 := person{
		First: "Farid",
		Last:  "Wajdi",
		Item:  []string{"1", "2", "3"},
	}
	d, err := tojson(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(d))
}
func tojson(a interface{}) ([]byte, error) {
	val, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("Terdapat Error pada %v", err)
		// hampir sama dengn dibawah
		return []byte{}, errors.New(fmt.Sprintf("Terdapat Error di %v", err))
	}
	return val, err
}
