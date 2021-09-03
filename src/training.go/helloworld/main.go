package main

import (
	"fmt"
)

func main() {
	weekDay := 6 // 1 == Monday, 7 == Sunday
	fmt.Printf("Day of the week=%d. What's for today?\n", weekDay)

	switch weekDay {
	case 1:
		fmt.Println("Beginning of the week, let's go to work!")
	case 3:
		fmt.Println("Wednesday, the half is done!")
	case 6, 7:
		fmt.Println("It's the week-end!")
	default:
		fmt.Println("Nothing special about this day...")
	}

	hour := 20
	switch {
	case hour < 12:
		fmt.Println("It's the morning!")
	case hour >= 12 && hour < 18:
		fmt.Println("It's the afternoon")
	default:
		fmt.Println("It's the evening")
	}
}
