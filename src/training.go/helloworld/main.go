package main

import (
	"fmt"
	"strings"

	"helloworld.com/greetings/src/training.go/helloworld/input"
)

func main() {
	input.Mouse()
	input.TotoMyFunc()
	fmt.Println(strings.ToUpper("Hello Gophers!"))
}
