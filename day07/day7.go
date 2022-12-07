package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	path := make([]string, 0)
	dirSizes := map[string]int{}

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		// caught command
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, line[2])
				}
			}
		} else {
			if line[0] != "dir" {
				currentDir := ""
				sizeInt, _ := strconv.Atoi(line[0])
				for i := 0; i < len(path); i++ {
					currentDir += path[i]
					value, exists := dirSizes[currentDir]
					if exists {
						dirSizes[currentDir] = value + sizeInt
					} else {
						dirSizes[currentDir] = sizeInt
					}
					if i > 0 {
						currentDir += "/"
					}
				}

			}
		}
	}

	solution := 0
	for _, value := range dirSizes {
		if value < 100000 {
			solution += value
		}
	}

	fmt.Printf("Solution for Day 7 Part 1: %d\n", solution)

}

func Part2() {
	f, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	path := make([]string, 0)
	dirSizes := map[string]int{}

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		// caught command
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, line[2])
				}
			}
		} else {
			if line[0] != "dir" {
				currentDir := ""
				sizeInt, _ := strconv.Atoi(line[0])
				for i := 0; i < len(path); i++ {
					currentDir += path[i]
					value, exists := dirSizes[currentDir]
					if exists {
						dirSizes[currentDir] = value + sizeInt
					} else {
						dirSizes[currentDir] = sizeInt
					}
					if i > 0 {
						currentDir += "/"
					}
				}

			}
		}
	}

	totalUsedSpace := dirSizes["/"]
	const FULL = 70000000
	unused := FULL - totalUsedSpace
	needed := 30000000 - unused
	solution := totalUsedSpace

	for _, value := range dirSizes {
		if value >= needed && value <= solution {
			solution = value
		}
	}

	fmt.Printf("Solution for Day 7 Part 2: %d\n", solution)

}
