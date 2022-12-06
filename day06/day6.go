package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solve(windowSize int) int {
	// open file
	f, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), "")
		starting := 0
		ending := windowSize
		for ending <= len(line) {
			chars := map[string]bool{}
			for _, value := range line[starting:ending] {
				_, exists := chars[value]
				if !exists {
					chars[value] = true
				}
			}
			if len(chars) == windowSize {
				return ending
			}
			starting++
			ending++
		}
	}
	return -1
}

func Part1() {
	fmt.Printf("Solution for Day 6 Part 1: %d\n", solve(4))
}

func Part2() {
	fmt.Printf("Solution for Day 6 Part 2: %d\n", solve(14))
}
