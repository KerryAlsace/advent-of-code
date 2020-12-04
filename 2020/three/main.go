package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// sample_input ans part 1 is 7; part 2 is 336;
// sample_input_two ans part 1 is 4;

// guessed 95, guessed 72, guessed 100;
// input ans part 1 is 216; part 2 is 6708199680;

func main() {
	i := getInput("input.txt")
	// partOneAnswer := partOne(i)
	partTwoAnswer := partTwo(i)

	// fmt.Println(partOneAnswer)
	fmt.Println(partTwoAnswer)
}

func getInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

// func partOne(i []string) int {
// 	spacesRight := 0
// 	totalTrees := 0

// 	for _, row := range i {
// 		if len(row) <= spacesRight {
// 			spacesRight = (spacesRight - (len(row) - 1)) - 1
// 		}

// 		x := row[spacesRight]

// 		if string(x) == "#" {
// 			totalTrees++
// 		}

// 		spacesRight += 3
// 	}

// 	return totalTrees
// }

func toboggan(right, down int, input []string) int {
	spacesRight := 0
	totalTrees := 0

	for i, row := range input {
		if down == 2 && i%2 != 0 {
			continue
		}

		if len(row) <= spacesRight {
			spacesRight = (spacesRight - (len(row) - 1)) - 1
		}

		x := row[spacesRight]

		if string(x) == "#" {
			totalTrees++
		}

		spacesRight += right
	}

	return totalTrees
}

func partTwo(i []string) int {
	return toboggan(1, 1, i) * toboggan(3, 1, i) * toboggan(5, 1, i) * toboggan(7, 1, i) * toboggan(1, 2, i)
}
