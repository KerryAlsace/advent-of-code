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
	partTwoAnswer := partTwo(i)

	fmt.Println(partOneAnswer)
	fmt.Println(partTwoAnswer)
}

func getInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(i []string) int {
	return iterator(i, passwordIsValidPartOne)
}

func partTwo(i []string) int {
	return iterator(i, passwordIsValidPartTwo)
}

func iterator(i []string, checker func(min, max int, letter, password string) bool) int {
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

		if checker(min, max, strings.TrimRight(parts[1], ":"), parts[2]) {
			correctCount++
		}
	}

	return correctCount
}

func passwordIsValidPartOne(min, max int, letter, password string) bool {
	var letterCount int
	for _, l := range password {
		if string(l) == letter {
			letterCount++
		}
	}

	return min <= letterCount && letterCount <= max
}

func passwordIsValidPartTwo(min, max int, letter, password string) bool {
	var letterCount int

	firstIndex := min - 1
	secondIndex := max - 1

	if string(password[firstIndex]) == letter {
		letterCount++
	}

	if string(password[secondIndex]) == letter {
		letterCount++
	}

	return letterCount == 1
}
