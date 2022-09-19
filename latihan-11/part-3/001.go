package main

import "fmt"

type myerror struct {
	info string
}

func (x myerror) Error() string {
	return fmt.Sprintf("Terdapat Error pada %v", x.info)
}
func main() {
	p1 := myerror{
		info: "jaringan",
	}
	foo(p1)

}
func foo(e error) {
	fmt.Println("foo - ", e)
	// convertions
	fmt.Println("foo - ", e.(myerror).info)
}
