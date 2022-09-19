package main

import (
	"errors"
	"fmt"
	"log"
)

type myerr struct {
	lat  string
	long string
	err  error
}

func (x myerr) Error() string {
	return fmt.Sprintf("Terdapat Error pada %v, %v, %v", x.lat, x.long, x.err)
}
func main() {
	i, err := scp(110)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
func scp(x int) (int, error) {
	if x > 10 {
		e := errors.New("parah")
		return 10, myerr{"Satu", "dua", e}
	}
	return x, nil

}
