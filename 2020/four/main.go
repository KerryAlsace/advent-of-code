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
	firstLineOfLastInput := 0
	lastLineOfLastInput := 0
	for index, line := range i {
		fmt.Printf("index [%v] line [%s]\n", index, line)
		if line == "" || index == len(i)-1 {
			if line == "" {
				lastLineOfLastInput = index - 1
			} else if index == len(i)-1 {
				lastLineOfLastInput = index
			}

			fmt.Printf("lastLineOfLastInput [%v]\n", lastLineOfLastInput)

			var sb strings.Builder
			for n := 0; n <= lastLineOfLastInput-firstLineOfLastInput; n++ {
				fmt.Printf("n = [%v] lastlineof [%v], firstline of [%v] lastLineOfLastInput-firstLineOfLastInput = [%v]\n", n, lastLineOfLastInput, firstLineOfLastInput, lastLineOfLastInput-firstLineOfLastInput)
				sb.WriteString(i[firstLineOfLastInput+n])
				sb.WriteString("")
			}

			firstLineOfLastInput = index + 1
			newInput = append(newInput, sb.String())
		}
	}

	return newInput
}

func partOne(i []string) int {
	for _, r := range i {
		fmt.Println(r)
	}
	return 0
}

// func partTwo(i []string) int {
// 	return 0
// }
