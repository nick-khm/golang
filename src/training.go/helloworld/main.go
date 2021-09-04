package main

import (
	"fmt"
	"strings"
)

func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {

	lowerName, len := ToLowerStr("ALICE")
	fmt.Printf("%s (len=%d)\n", lowerName, len)

	lowerName, len = ToLowerStr("JeelyBom")
	fmt.Printf("%s (len=%d)\n", lowerName, len)

	bobName, len := ToLowerStr("Bob")
	fmt.Printf("%s (len=%d)\n", bobName, len)

	_, len = ToLowerStr("Bob")
	fmt.Printf("String length=%d)\n", len)

	stringName, _ := ToLowerStr("Bob")
	fmt.Printf("String name=%s\n", stringName)

}
