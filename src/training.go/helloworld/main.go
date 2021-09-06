package main

import (
	"fmt"
)

func main() {
	names := []string{"Bob", "Alice", "Bobette", "John"}
	for i, name := range names {
		fmt.Printf("Name=%s, index=%d\n", name, i)
	}
}
