package main

import (
	"bufio"
	"fmt"
	"os"
)

const SAMPLE_TARGET_P1_01 = 5
const SAMPLE_TARGET_P1_02 = 6
const SAMPLE_TARGET_P1_03 = 10
const SAMPLE_TARGET_P1_04 = 11
const SAMPLE_TARGET_P2_01 = 23
const SAMPLE_TARGET_P2_02 = 23
const SAMPLE_TARGET_P2_03 = 29
const SAMPLE_TARGET_P2_04 = 26

func main() {
	scanner := makeScanner("input.txt")
	resultP1 := calculate(scanner, 4)

	scanner = makeScanner("input.txt")
	resultP2 := calculate(scanner, 14)

	fmt.Printf("Part 1: First marker = %d\n", resultP1)
	fmt.Printf("Part 2: First marker = %d\n", resultP2)
}

func calculate(scanner *bufio.Scanner, windowSize int) int {
	result := 0

	scanner.Scan()
	line := scanner.Text()

	for offset := 0; offset < len(line) - windowSize; offset++ {
		// Get window
		window := line[offset:offset+windowSize]
		// fmt.Printf("Off %d, Window %s\n", offset, window)

		// Find dupes
		chars := make(map[rune]bool)

		for _, c := range window {
			chars[c] = true
		}

		if len(chars) == windowSize {
			return offset + windowSize
		}
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
