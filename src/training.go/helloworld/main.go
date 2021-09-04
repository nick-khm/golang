package main

import (
	"fmt"
)

func printInfoParams(name string, age int, email string) {
	fmt.Printf("Name=%s, age=%d, email=%s\n", name, age, email)
}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return sum
}

func sumNamedReturn2(x, y, z int) (sum int) {
	sum = x + y + z
	return
}

func main() {
	printInfoParams("Bob", 20, "bob@t.com")
	avgResult := avg(32, 34.3)
	fmt.Printf("AVG result=%f\n", avgResult)
	sum := sumNamedReturn(8, 32, 5)
	fmt.Printf("Sum result=%d\n", sum)
	sum2 := sumNamedReturn2(8, 32, 5)
	fmt.Printf("Sum result=%d\n", sum2)
}
