package main

import (
	"fmt"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(res, old) || strings.Contains(res, oldLower) {
		found = true
		occ += strings.Count(res, old)
		occ += strings.Count(res, oldLower)
		res = strings.Replace(res, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func main() {
	found, res, occ := ProcessLine("Go is the best language to learn. You should go to the official website.", "Go", "Python")
	fmt.Println(found, res, occ)
}
