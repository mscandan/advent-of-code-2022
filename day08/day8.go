package day08

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFileAndReturnGrid() [][]int {
	grid := make([][]int, 0)
	f, err := os.Open("day08/input.txt")
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
		row := make([]int, 0)
		line := strings.Split(scanner.Text(), "")
		for _, value := range line {
			intValue, _ := strconv.Atoi(value)
			row = append(row, intValue)
		}
		grid = append(grid, row)
	}
	return grid
}

func Part1() {
	grid := readFileAndReturnGrid()
	solution := 0

	rowLength := len(grid)
	columnLength := len(grid[0])

	for row := 0; row < rowLength; row++ {
		for column := 0; column < columnLength; column++ {
			if row == 0 || row == rowLength-1 || column == 0 || column == columnLength-1 {
				solution++
			} else {
				curr := grid[row][column]
				visibilityArray := make([]bool, 0)
				// left check
				isVisibleFromLeft := true
				for k := column - 1; k >= 0; k-- {
					if grid[row][k] >= curr {
						isVisibleFromLeft = false
					}
				}
				if isVisibleFromLeft {
					visibilityArray = append(visibilityArray, true)
				} else {
					visibilityArray = append(visibilityArray, false)
				}
				// right check
				isVisibleFromRight := true
				for k := column + 1; k < columnLength; k++ {
					if grid[row][k] >= curr {
						isVisibleFromRight = false
					}
				}
				if isVisibleFromRight {
					visibilityArray = append(visibilityArray, true)
				} else {
					visibilityArray = append(visibilityArray, false)
				}
				// top check
				isVisibleFromTop := true
				for k := row - 1; k >= 0; k-- {
					if grid[k][column] >= curr {
						isVisibleFromTop = false
					}
				}
				if isVisibleFromTop {
					visibilityArray = append(visibilityArray, true)
				} else {
					visibilityArray = append(visibilityArray, false)
				}
				// bottom check
				isVisibleFromBottom := true
				for k := row + 1; k < rowLength; k++ {
					if grid[k][column] >= curr {
						isVisibleFromBottom = false
					}
				}
				if isVisibleFromBottom {
					visibilityArray = append(visibilityArray, true)
				} else {
					visibilityArray = append(visibilityArray, false)
				}

				for _, value := range visibilityArray {
					if value {
						solution++
						break
					}
				}
			}
		}
	}

	fmt.Printf("Solution for Day 8 Part 1:  %d\n", solution)

}

func Part2() {
	grid := readFileAndReturnGrid()
	solution := 0.0

	for rowIdx, row := range grid {
		for columnIdx := range row {
			curr := grid[rowIdx][columnIdx]
			distanceLeft, distanceRight, distanceTop, distanceBottom := 0, 0, 0, 0
			if rowIdx != 0 && rowIdx != len(grid)-1 && columnIdx != 0 && columnIdx != len(row)-1 {
				// left distance
				for i := columnIdx - 1; i >= 0; i-- {
					if grid[rowIdx][i] < curr {
						distanceLeft++
					} else {
						distanceLeft++
						break
					}
				}
				// right distance
				for i := columnIdx + 1; i < len(row); i++ {
					if grid[rowIdx][i] < curr {
						distanceRight++
					} else {
						distanceRight++
						break
					}
				}
				// top distance
				for i := rowIdx - 1; i >= 0; i-- {
					if grid[i][columnIdx] < curr {
						distanceTop++
					} else {
						distanceTop++
						break
					}
				}
				// bottom distance
				for i := rowIdx + 1; i < len(grid); i++ {
					if grid[i][columnIdx] < curr {
						distanceBottom++
					} else {
						distanceBottom++
						break
					}
				}
				score := distanceLeft * distanceRight * distanceTop * distanceBottom
				solution = math.Max(solution, float64(score))
			}
		}

	}

	fmt.Printf("Solution for Day 8 Part 2:  %d\n", int(solution))

}
