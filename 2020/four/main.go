package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i := getInput("sample_input.txt")
	partOneAnswer := partOne(formatInput(i))
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

func formatInput(i []string) []string {
	var newInput []string
	currentIndex := 0
	lastLineOfLastInput := 0
	for index, line := range i {
		fmt.Println("here's the line")
		fmt.Println(line)
		if line == "" {
			fmt.Println("line == ''")
			fmt.Printf("index is %v line is %s\n", index, line)
			lastLineOfLastInput = index - 1

			var sb strings.Builder
			for n := 0; n <= lastLineOfLastInput-currentIndex; n++ {
				fmt.Printf("current index %v\n", currentIndex)
				sb.WriteString(i[currentIndex])
				currentIndex++
			}

			newInput = append(newInput, sb.String())
		}
	}

	return newInput
}

func partOne(i []string) int {
	// for _, r := range i {
	// 	// fmt.Println(r)
	// }
	return 0
}

// func partTwo(i []string) int {
// 	return 0
// }
