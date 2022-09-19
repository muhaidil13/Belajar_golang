package main

import (
	"os"
)

// printing and logging
func main() {
	// printting
	// f, err := os.Open("Test")
	// if err != nil {
	// 	// dapat di kostum textnya
	// 	fmt.Println("Terdapat Error ", err)
	// }
	// defer f.Close()
	//
	//
	// logging-set-output
	// f, err := os.Create("log.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	// f, err = os.Open("op")
	// if err != nil {
	// 	log.Println("Terdapat error di ", err)
	// }
	// fmt.Println("cek l.txt")
	//
	//
	// Fatal
	// hampir sama dengan println tapi memanggil fungsi os exit
	// jika error maka code akan berhenti seperti defer
	// f, err := os.Open("test")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f.Close()

	//
	// log panic
	// hampir sama dengan fmt.println
	// menghentikan semua termasuk goruntime
	// f, err := os.Open("test")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// f.Close()

	// panic
	// mengentikan exekusi normal yang dijalankan sekarang
	_, err := os.Open("test")
	if err != nil {
		panic(err)
	}

}
