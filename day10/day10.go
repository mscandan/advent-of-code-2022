package day10

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	X := 1
	cycleCount := 0
	cycleValues := map[int]int{}

	f, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "noop" {
			cycleCount++
			cycleValues[cycleCount] = cycleCount * X
		} else {
			cycleCount++
			cycleValues[cycleCount] = cycleCount * X
			cycleCount++
			cycleValues[cycleCount] = cycleCount * X
			amount, _ := strconv.Atoi(line[1])
			X += amount
		}
	}

	fmt.Printf("Solution for Day 10 Part 1: %d\n", cycleValues[20]+cycleValues[60]+cycleValues[100]+cycleValues[140]+cycleValues[180]+cycleValues[220])
}

func Part2() {
	X := 1
	cycleCount := 0
	cycleValues := map[int]int{}

	f, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "noop" {
			cycleCount++
			cycleValues[cycleCount] = X
		} else {
			cycleCount++
			cycleValues[cycleCount] = X
			cycleCount++
			cycleValues[cycleCount] = X
			amount, _ := strconv.Atoi(line[1])
			X += amount
		}
	}

	fmt.Println("Solution for Day 10 Part2: ")
	for i := 0; i < 6; i++ {
		row := ""
		for j := 0; j < 40; j++ {
			count := i*40 + j + 1
			if math.Abs(float64(cycleValues[count])-float64(j)) <= 1 {
				row += "#"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}
