package main

import (
	"fmt"
	"math"
)

// 4
type person struct {
	first string
	last  string
	age   int
}

// 5
type bulat struct {
	radius float64
}
type kotak struct {
	length float64
}
type tampil interface {
	area() float64
}

func main() {
	// 3
	defer intro()

	// 1
	fmt.Println(foo(19))
	fmt.Println(bar())

	// 2
	x := []int{10, 20, 30, 40}
	fmt.Println(sum(x...))
	fmt.Println(sum2(x))

	// 4
	uni := person{
		first: "uni",
		last:  "saputra",
		age:   21,
	}
	uni.speak()

	// 5
	xb := bulat{
		radius: 20.3,
	}
	xk := kotak{
		length: 12.5,
	}
	info(xb, "Jumlah")
	info(xk, "adalah")

	// 6
	func() {
		fmt.Println("Ini adalah anynymouse function")
	}()
	func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	}()

	// 7
	dq := func() {
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	}
	dq()

	// 8
	ac := test()
	fmt.Println(ac())

	// 9
	xxa := []int{1, 2, 3, 4, 5, 6}
	ada := display(tambah, xxa)
	fmt.Println(ada)

	// 10
	cl := closures()
	cls := closures1()
	fmt.Println(cl())
	fmt.Println(cls)

	// 11

}

// 1
func foo(x int) int {
	total := x + x
	return total
}
func bar() (int, string) {
	return 10, "ini adalah string"
}

// 2
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func sum2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// 3
func intro() {
	defer func() {
		fmt.Println("=========END======")
	}()
	fmt.Println("========================")
}

// 4
func (x person) speak() {
	fmt.Println("nama depan saya adalah ", x.first)
	fmt.Println("nama belakang saya adalah ", x.last)
}

// 5
// method
func (x bulat) area() float64 {
	return math.Pi * x.radius * x.radius
}

// method
func (x kotak) area() float64 {
	return x.length * x.length
}

func info(x tampil, y string) {
	fmt.Println(x.area(), y)
}

// 8
func test() func() int {
	data := 10 * 10
	return func() int {
		return data
	}
}

// 9
func tambah(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func display(f func(x []int) int, xi []int) int {
	n := f(xi)
	n++
	n++
	n++
	n++
	n++
	return n
}

// 10
func closures() func() int {
	var data int
	return func() int {
		data++
		return data
	}
}
func closures1() int {
	var data int
	return func() int {
		data++
		return data
	}()
}
