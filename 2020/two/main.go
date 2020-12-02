package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	i := getInput("input.txt")
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
		minMax := strings.Split(parts[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			panic("invalid input")
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			panic("invalid input")
		}

		if passwordIsValid(min, max, strings.TrimRight(parts[1], ":"), parts[2]) {
			correctCount++
		}
	}

	return correctCount
}

func passwordIsValid(min, max int, letter, password string) bool {
	var letterCount int
	for _, l := range password {
		if string(l) == letter {
			letterCount++
		}
	}

	if min <= letterCount && letterCount <= max {
		return true
	}

	return false
}

// func partTwo(i []string) int {
// 	return 0
// }
