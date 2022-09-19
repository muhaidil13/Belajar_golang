package main

import (
	"errors"
	"fmt"
	"log"
)

// error with info
var errf = errors.New("\nTest error")

func main() {
	// fmt.Printf("%T", errf)
	// a, err := mod(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a)
	_, err := smp(1)
	if err != nil {
		log.Fatal(err)
	}
}

//	func mod(x int) (int, error) {
//		if x > 5 {
//			return 5, errf
//		} else if x == 0 {
//			return 0, errors.New("\nNilai Tidak Boleh Nol")
//		}
//		return x, nil
//	}
func smp(f int) (int, error) {
	if f == 0 {
		return 5, fmt.Errorf("error %v", f)
	} else if f < 2 {
		errorcode := fmt.Errorf("errpr value %v ", f)
		return f, errorcode
	}
	return f, nil
}
