package main

import (
	"fmt"
	"log"
)

type myerr struct {
	name string
	id   string
	err  error
}

func (s myerr) Error() string {
	return fmt.Sprintf("terdapat error ada %v, %v, %v", s.name, s.id, s.err)
}

func main() {
	// costume error
	_, err := sqrt(1)
	if err != nil {
		log.Fatal(err)
	}

}
func sqrt(x int) (int, error) {
	if x < 10 {
		nm := fmt.Errorf("error code ")
		return 0, myerr{"terlalu tinggi", "1", nm}
	}
	return x, nil
}
