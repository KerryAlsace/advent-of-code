package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i := getInput("sample_input.txt")
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
	return 0
}

func partTwo(i []string) int {
	return 0
}
