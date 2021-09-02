package main

import (
	"fmt"
)

var pi = 3.14

func main() {
	var age int = 20
	fmt.Println(age)

	var weight, height int = 80, 190
	fmt.Println(weight, height)

	sex, politic := "male", "respublican"
	fmt.Println(sex, politic)

	kids, car := 0, "pontiac"
	fmt.Println(kids, car)

	fmt.Println(pi)
}
