package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// sample_input ans part 1 is 7; part 2 is
// sample_input_two ans part 1 is 1;

// guessed 95, guessed 72, guessed 100

func main() {
	i := getInput("sample_input_two.txt")
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
	spacesRight := 0
	totalTrees := 0

	for _, row := range i {
		if len(row) <= spacesRight {
			spacesRight = (spacesRight - (len(row) - 1)) - 1
			continue
		}

		x := row[spacesRight]

		if string(x) == "#" {
			totalTrees++
		}

		spacesRight += 3
	}

	return totalTrees
}

// func partTwo(i []string) int {
// 	return 0
// }
