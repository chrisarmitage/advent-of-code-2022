package main

import (
	"bufio"
	"fmt"
	"os"
)

const SAMPLE_TARGET_P1 = 2
const SAMPLE_TARGET_P2 = 4

type Assignment struct {
	P1Start int
	P1End int
	P2Start int
	P2End int
}

func main() {
	scanner := makeScanner("input.txt")
	resultP1 := calculatePart1(scanner)
	
	scanner = makeScanner("input.txt")
	resultP2 := calculatePart2(scanner)

	fmt.Printf("Part 1: Fully contained assignments = %d\n", resultP1)
	fmt.Printf("Part 2: Partially contained assignments = %d\n", resultP2)
}

func calculatePart1(scanner *bufio.Scanner) int {
	result := 0

	
	for scanner.Scan() {
        line := scanner.Text()

		var a Assignment

		fmt.Sscanf(
			line, "%d-%d,%d-%d",
			&a.P1Start, &a.P1End, &a.P2Start, &a.P2End)
		
			
		// fmt.Printf("%v", a)
		if (a.P2Start >= a.P1Start && a.P2End <= a.P1End) ||
		(a.P2Start <= a.P1Start && a.P2End >= a.P1End) {
			result++
			// fmt.Printf(" - Overlap")
		} 

		// fmt.Print("\n")
	}

	return result
}

func calculatePart2(scanner *bufio.Scanner) int {
	result := 0
	
	for scanner.Scan() {
        line := scanner.Text()

		var a Assignment

		fmt.Sscanf(
			line, "%d-%d,%d-%d",
			&a.P1Start, &a.P1End, &a.P2Start, &a.P2End)
		
			
		// fmt.Printf("%v", a)
		if (a.P2Start >= a.P1Start && a.P2Start <= a.P1End) ||
		(a.P2End >= a.P1Start && a.P2End <= a.P1End) ||
		(a.P1Start >= a.P2Start && a.P1Start <= a.P2End) ||
		(a.P1End >= a.P2Start && a.P1End <= a.P2End){
			result++
			// fmt.Printf(" - Overlap")
		} 

		// fmt.Print("\n")
	}

	return result
}

func makeScanner(filename string) *bufio.Scanner {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}