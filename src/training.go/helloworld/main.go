package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(dat) == 0 {
		// return "", errors.New("Empty content")
		return "", fmt.Errorf("empty content (filename=%v)", filename)
	}

	return string(dat), nil
}

func main() {
	dat, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(dat)
}
