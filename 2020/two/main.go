package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i := getInput("sample_input.txt")
	partOneAnswer := partOne(i)
	// partTwoAnswer := partTwo(i)

	fmt.Println(partOneAnswer)
	// fmt.Println(partTwoAnswer)
}

func getInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(i []string) int {
	var correctCount int

	for _, s := range i {
		parts := strings.Split(s, " ")
		letter := strings.TrimRight(parts[1], ":")
		minMax := strings.Split(parts[0], "-")
		min := minMax[0]
		max := minMax[1]
		password := parts[2]

		fmt.Printf("%s\n%s\n%s\n%s\n", letter, min, max, password)
	}

	return correctCount
}

// func partTwo(i []string) int {
// 	return 0
// }
