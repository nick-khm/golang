package main

import (
	"fmt"
)

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

func main() {
	start()
	defer finish() // LIFO = Last In First Out
	// when closing the CURRENT FUNCTION it will stock func to finish
	// finish()
	// Goodbye Bob
	// Goodbye Alice
	// Goodbye Bobette
	// Goodbye John

	names := []string{"Bob", "Alice", "Bobette", "John"}
	for _, n := range names {
		fmt.Printf("Hello %v\n", n)
		defer fmt.Printf("Goodbye %v\n", n)
	}

	/*
		Start
		Hello Bob
		Hello Alice
		Hello Bobette
		Hello John
		Goodbye John
		Goodbye Bobette
		Goodbye Alice
		Goodbye Bob
		Finish
	*/
}
