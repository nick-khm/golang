package main

import (
	"fmt"

	"helloworld.com/greetings/src/training.go/helloworld/data"
)

func main() {
	fmt.Println(data.Name, data.Age)

	if count := 12; count > 10 {
		fmt.Printf("Counter is enough, = %d\n", count)
	} else {
		fmt.Printf("Counter is not enough, = %d\n", count)
	}
}
