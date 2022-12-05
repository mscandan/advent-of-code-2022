package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() {

	// open file
	f, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	priority := 0

	for scanner.Scan() {
		// do something with a line
		seenChars := map[string]int{}
		priorityCalculted := map[string]int{}
		allChars := strings.Split(scanner.Text(), "")
		for i := 0; i < (len(allChars) / 2); i++ {
			seenChars[allChars[i]] = 1
		}
		for i := len(allChars) / 2; i < len(allChars); i++ {
			_, exists := seenChars[allChars[i]]
			if exists {
				_, isAlreadyCalculated := priorityCalculted[allChars[i]]
				if !isAlreadyCalculated {
					priority += calculatePriority(allChars[i])
					priorityCalculted[allChars[i]] = 1
				}
			}
		}
	}

	fmt.Printf("Day3 Part1 Solution: %d\n", priority)
}

func Part2() {

	// open file
	f, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	priority := 0

	lineCount := 0
	group := make([]string, 0)
	for scanner.Scan() {
		// do something with a line
		lineCount++
		line := scanner.Text()
		group = append(group, line)
		if lineCount%3 == 0 {
			first := group[lineCount-3]
			second := group[lineCount-2]
			third := group[lineCount-1]
			seenChars := map[string]bool{}
			for i := 0; i < len(first); i++ {
				_, isAlreadyCalculated := seenChars[string(first[i])]
				if !isAlreadyCalculated {
					if contains(second, first[i]) && contains(third, first[i]) {
						priority += calculatePriority(string(first[i]))
						seenChars[string(first[i])] = true
					}
				}
			}
		}

	}

	fmt.Printf("Day3 Part2 Solution: %d\n", priority)
}

func contains(arr string, target byte) bool {
	for _, value := range arr {
		if value == rune(target) {
			return true
		}
	}
	return false
}

func calculatePriority(char string) int {
	charCode := int(char[0])
	if 65 <= charCode && charCode <= 90 {
		return charCode - 38
	}
	return charCode - 96
}
