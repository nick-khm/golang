package main

import (
	"fmt"
	"strings"
)

func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {
	nums := make([]int, 2, 3)
	nums[0] = 23
	nums[1] = 64
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, 8)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	// will double a capacity
	nums = append(nums, 99)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	// dynamic slice
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("letters %v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Println("-----------------------------")

	// sub slice
	sub1 := letters[0:2] // or letters [:2]
	sub2 := letters[2:]  // will take untile the end
	fmt.Printf("sub1 %v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2 %v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("-----------------------------")

	// slice changeing will impact the source array
	sub2[0] = "YOLO"
	fmt.Printf("sub1 %v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2 %v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("letters %v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Println("-----------------------------")

	// copy a part of an array, not referenced
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[1] = "WIN"
	fmt.Printf("sub1 %v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2 %v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("subCopy %v (len=%d, cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("letters %v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Println("-----------------------------")
}
