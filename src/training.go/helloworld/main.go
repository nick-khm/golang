package main

import (
	"fmt"
	"strings"
)

func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {
	i := 0
	for {
		i++
		if i%2 != 0 {
			fmt.Println("Odd looping")
			continue
		}
		fmt.Println("Looping")
		if i >= 10 {
			fmt.Println("Exit looping")
			break
		}
	}
}
