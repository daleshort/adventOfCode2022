package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	filepath := "day6.txt"
	input, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	inputStr := string(input)

	inputSplit := strings.Split(inputStr, "\n")

	for _, s := range inputSplit {

		//bvwbjplbgvbhsrlpgdmjqwftvncz
		// |  |
		// front pointer starts at 0
		// back pointer starts at front pointer + 3
		// grab all the characters between front pointer and back pointer
		// check to see if the characters are all unique
		// if they are all unique we return the location of front pointer-1
		// if back pointer is greater than the length of string we have error

		front := 0
		back := 13

		for back < len(s) {
			if isUniqueSubstring(s[front : back+1]) {
				fmt.Println("start code found", back+1)
				break
			}
			front++
			back++

		}
		fmt.Println("_____")
	}
}

func isUniqueSubstring(s string) bool {
	searchMap := make(map[rune]struct{})

	for _, c := range s {
		if _, ok := searchMap[c]; ok {
			return false
		}
		searchMap[c] = struct{}{}
	}
	return true

}
