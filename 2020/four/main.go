package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// sample input ans part 1 is 2;

func main() {
	i := getInput("input.txt")
	// partOneAnswer := partOne(formatInput(i))
	partTwoAnswer := partTwo(formatInput(i))

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

func formatInput(i []string) []string {
	var newInput []string
	firstLineOfLastInput := 0
	lastLineOfLastInput := 0
	for index, line := range i {
		if line == "" || index == len(i)-1 {
			if line == "" {
				lastLineOfLastInput = index - 1
			} else if index == len(i)-1 {
				lastLineOfLastInput = index
			}

			var sb strings.Builder
			for n := 0; n <= lastLineOfLastInput-firstLineOfLastInput; n++ {
				sb.WriteString(i[firstLineOfLastInput+n])
				sb.WriteString(" ")
			}

			firstLineOfLastInput = index + 1
			newInput = append(newInput, sb.String())
		}
	}

	return newInput
}

func partOne(i []string) int {
	validPassports := 0

	for _, r := range i {
		comp := make(map[string]int)

		n := strings.Split(r, " ")
		for _, l := range n {
			p := strings.Split(l, ":")
			comp[p[0]] = 1
		}

		if isValid(comp) {
			validPassports++
		}
	}

	return validPassports
}

func isValid(i map[string]int) bool {
	return i["ecl"] == 1 && i["pid"] == 1 && i["eyr"] == 1 && i["hcl"] == 1 && i["byr"] == 1 && i["iyr"] == 1 && i["hgt"] == 1
}

// func partTwo(i []string) int {
// 	return 0
// }
