package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i := getInput("sample_input.txt")
	partOneAnswer := partOne(i)

	fmt.Println(partOneAnswer)
}

func getInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(i []string) string {
	for _, s := range i {
		fmt.Println(s)
	}
	return ""
}
