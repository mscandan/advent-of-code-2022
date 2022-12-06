package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Part1() {

	// open file
	f, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	isManipulatePart := false

	stacks := make([][]string, 9)
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()

		// assign to arrays part
		if line == "" {
			isManipulatePart = true
			for _, stack := range stacks {
				sort.Slice(stack, func(i, j int) bool {
					return i > j
				})
			}
			continue
		}

		if !isManipulatePart {
			splittedLine := strings.Split(line, "")
			for idx, value := range splittedLine {
				if value != " " && value != "[" && value != "]" && !is_numeric(value) {
					targetStackIdx := (idx - 1) / 4
					stacks[targetStackIdx] = append(stacks[targetStackIdx], value)
				}
			}
		} else {
			splittedLine := strings.Split(line, " ")
			moveCount, _ := strconv.Atoi(splittedLine[1])
			fromIdx, _ := strconv.Atoi(splittedLine[3])
			toIdx, _ := strconv.Atoi(splittedLine[5])
			fromIdx--
			toIdx--
			for i := 0; i < moveCount; i++ {
				poppedVal := stacks[fromIdx][len(stacks[fromIdx])-1]
				stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-1]
				stacks[toIdx] = append(stacks[toIdx], poppedVal)
			}
		}
	}

	fmt.Print("Solution for Day 5 Part 1: ")
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Print("\n")
}

func Part2() {

	// open file
	f, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	isManipulatePart := false

	stacks := make([][]string, 9)
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()

		// assign to arrays part
		if line == "" {
			isManipulatePart = true
			for _, stack := range stacks {
				sort.Slice(stack, func(i, j int) bool {
					return i > j
				})
			}
			continue
		}

		if !isManipulatePart {
			splittedLine := strings.Split(line, "")
			for idx, value := range splittedLine {
				if value != " " && value != "[" && value != "]" && !is_numeric(value) {
					targetStackIdx := (idx - 1) / 4
					stacks[targetStackIdx] = append(stacks[targetStackIdx], value)
				}
			}
		} else {
			splittedLine := strings.Split(line, " ")
			moveCount, _ := strconv.Atoi(splittedLine[1])
			fromIdx, _ := strconv.Atoi(splittedLine[3])
			toIdx, _ := strconv.Atoi(splittedLine[5])
			fromIdx--
			toIdx--
			itemsToMove := make([]string, 0)
			for i := 0; i < moveCount; i++ {
				poppedVal := stacks[fromIdx][len(stacks[fromIdx])-1]
				stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-1]
				itemsToMove = append(itemsToMove, poppedVal)
			}
			sort.Slice(itemsToMove, func(i, j int) bool {
				return i > j
			})
			stacks[toIdx] = append(stacks[toIdx], itemsToMove...)
		}
	}

	fmt.Print("Solution for Day 5 Part 2: ")
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Print("\n")
}

func is_numeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
}
