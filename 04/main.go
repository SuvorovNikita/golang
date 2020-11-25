package main

import "fmt"

// const applicationName = "App Const"
const (
	a = iota
	b = iota
	c = iota
	d = iota + 5
)

func main() {
	// const number = 100

	// fmt.Println(applicationName)
	// fmt.Println(number)

	// var numberTwo int = 10

	// result := number + numberTwo
	// fmt.Println(result)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	x, y, z := TypesGo()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func TypesGo() (int, bool, string) {
	return 100, false, "StringType"
}
