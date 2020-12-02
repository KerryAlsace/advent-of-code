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
	for _, x := range i {
		a, err := strconv.Atoi(x)
		if err != nil {
			panic("invalid input")
		}
		for _, y := range i {
			b, err := strconv.Atoi(y)
			if err != nil {
				panic("invalid input")
			}

			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func partTwo(i []string) int {
	for _, x := range i {
		a, err := strconv.Atoi(x)
		if err != nil {
			panic("invalid input")
		}
		for _, y := range i {
			b, err := strconv.Atoi(y)
			if err != nil {
				panic("invalid input")
			}
			for _, z := range i {
				c, err := strconv.Atoi(z)
				if err != nil {
					panic("invalid input")
				}

				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}
