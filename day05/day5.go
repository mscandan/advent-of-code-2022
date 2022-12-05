package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	//	maxnum := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		splittedLine := strings.Split(scanner.Text(), " ")
		fmt.Println(splittedLine)
	}

}
