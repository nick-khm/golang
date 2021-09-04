package main

import (
	"fmt"
)

func main() {
	var percentage float64 = 2.5

	percentage = percentage + 25

	fmt.Printf("Percentage %f%%\n", percentage)
	fmt.Printf("Int value %d%%\n", int(percentage))

	n := 42
	f := float64(n) + 0.42
	fmt.Printf("float=%f\n", f)
}
