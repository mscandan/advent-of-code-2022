package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	solution := 0

	f, err := os.Open("day4/input.txt")
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
		line := strings.Split(scanner.Text(), ",")
		firstPoints := strings.Split(line[0], "-")
		secondPoints := strings.Split(line[1], "-")
		starting1, _ := strconv.Atoi(firstPoints[0])
		ending1, _ := strconv.Atoi(firstPoints[1])

		starting2, _ := strconv.Atoi(secondPoints[0])
		ending2, _ := strconv.Atoi(secondPoints[1])

		if (starting2 <= starting1 && ending2 >= ending1) || (starting2 >= starting1 && ending2 <= ending1) {
			solution++
		}

	}
	fmt.Printf("Solution for Day 4 Part 1: %d\n", solution)
}

func Part2() {
	solution := 0

	f, err := os.Open("day4/input.txt")
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
		line := strings.Split(scanner.Text(), ",")
		firstPoints := strings.Split(line[0], "-")
		secondPoints := strings.Split(line[1], "-")
		starting1, _ := strconv.Atoi(firstPoints[0])
		ending1, _ := strconv.Atoi(firstPoints[1])
		starting2, _ := strconv.Atoi(secondPoints[0])
		ending2, _ := strconv.Atoi(secondPoints[1])

		if (ending1 >= starting2 && starting1 <= ending2) || (ending2 >= starting1 && ending2 <= ending1) {
			solution++
		}

	}
	fmt.Printf("Solution for Day 4 Part 2: %d\n", solution)
}
