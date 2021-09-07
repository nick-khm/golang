package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	dat, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot open (filename=%v)", src)
	}
	// fmt.Printf("Dat content: %v\n", string(dat))
	var replacedDat string = strings.Replace(string(dat), " Go.", "Python.", -1)
	replacedDat = strings.Replace(replacedDat, " Go ", " Python ", -1)
	replacedDat = strings.Replace(replacedDat, "Go ", "Python ", -1)
	fmt.Printf("Replaced content: \n\n %v\n", replacedDat)
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	i := 0
	for scanner.Scan() {
		i++
		t := scanner.Text()
		lines = append(lines, i)
		// fmt.Printf("Line %d : %v\n", i, t)
		cnt := strings.Count(t, "Go")
		// fmt.Printf("Count %v\n", cnt)
		occ += cnt
	}
	outputFile, fileCreateErr := os.Create("output.txt")
	if fileCreateErr != nil {
		return 0, nil, fmt.Errorf("cannot create file (filename=%v)", outputFile.Name())
	}
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	fmt.Fprintln(writer, replacedDat)
	return occ, lines, err
}

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	return
}

func main() {
	occ, lines, err := FindReplaceFile("wikigo.txt", "Go", "Python")
	if err != nil {
		fmt.Printf("Error occures while treatement of a file %v\n", err)
	}
	fmt.Printf("Occurances %v over %v lines\n", occ, len(lines))
}
