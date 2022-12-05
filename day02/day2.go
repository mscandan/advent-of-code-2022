package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() {
	choice1 := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	choice2 := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	// open file
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	myScore := 0

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		oppenentChoice := line[0]
		ourChoice := line[1]
		myScore += determineWinnerForPart1(choice1[oppenentChoice], choice2[ourChoice])
	}

	fmt.Printf("Day2 Part1 Solution: %d\n", myScore)

}

func Part2() {
	choice1 := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	// open file
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	myScore := 0

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		oppenentChoice := line[0]
		ourChoice := line[1]
		myScore += determineWinnerForPart2(choice1[oppenentChoice], ourChoice)
	}

	fmt.Printf("Day2 Part2 Solution: %d\n", myScore)
}

func determineWinnerForPart2(opponent string, our string) int {
	if opponent == "Rock" && our == "X" {
		return 3
	}
	if opponent == "Rock" && our == "Y" {
		return 4
	}
	if opponent == "Rock" && our == "Z" {
		return 8
	}
	if opponent == "Paper" && our == "X" {
		return 1
	}
	if opponent == "Paper" && our == "Y" {
		return 5
	}
	if opponent == "Paper" && our == "Z" {
		return 9
	}
	if opponent == "Scissors" && our == "X" {
		return 2
	}
	if opponent == "Scissors" && our == "Y" {
		return 6
	}

	return 7
}

func determineWinnerForPart1(opponent string, our string) int {
	if opponent == "Rock" && our == "Paper" {
		return 8
	}
	if opponent == "Rock" && our == "Rock" {
		return 4
	}
	if opponent == "Rock" && our == "Scissors" {
		return 3
	}
	if opponent == "Paper" && our == "Scissors" {
		return 9
	}
	if opponent == "Paper" && our == "Paper" {
		return 5
	}
	if opponent == "Paper" && our == "Rock" {
		return 1
	}
	if opponent == "Scissors" && our == "Rock" {
		return 7
	}
	if opponent == "Scissors" && our == "Scissors" {
		return 6
	}

	return 2
}
