package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Part1() {
	// open file
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	maxNum := 0
	currTotal := 0

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if len(line) > 0 {
			value, _ := strconv.Atoi(line)
			currTotal += value
		} else {
			if currTotal > maxNum {
				maxNum = currTotal
			}
			currTotal = 0
		}
	}

	fmt.Printf("Solution for day1 part1 = %d\n", maxNum)

}

func Part2() {
	// open file
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	allList := make([]int, 0)
	currTotal := 0

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if len(line) > 0 {
			value, _ := strconv.Atoi(line)
			currTotal += value
		} else {
			allList = append(allList, currTotal)
			currTotal = 0
		}
	}

	sort.Slice(allList, func(i, j int) bool {
		return allList[i] > allList[j]
	})

	fmt.Printf("solution for day1 part2 = %d", allList[0]+allList[1]+allList[2])

}
