package main

import (
	"fmt"
	"strings"
)

func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {

	var names [3]string
	fmt.Printf("names=%v (len=%v)\n", names, len(names))

	names[0] = "Bob"
	names[2] = "Alice"

	fmt.Printf("names=%v (len=%v)\n", names, len(names))

	odds := [4]int{44, 8, 32}
	fmt.Printf("odds=%v (len=%d)\n", odds, len(odds))
}
