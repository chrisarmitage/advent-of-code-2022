package main

import (
	"bufio"
	"fmt"
	"os"
)

const SAMPLE_TARGET_P1 = 157
const SAMPLE_TARGET_P2 = 70

func main() {
	scanner := makeScanner("input.txt")

	total := 0 

	resultP2 := calculatePart2(scanner)

	
	scanner = makeScanner("sample.txt")

	for scanner.Scan() {
        line := scanner.Text()

		compartmentSize := len([]rune(line)) / 2
		compartmentOne := line[0:compartmentSize]
		compartmentTwo := line[compartmentSize:]

		// Set up bitmasks
		var cpmOneBitmask uint64
		var cpmTwoBitmask uint64

		for _, item := range compartmentOne {
			// Convert each item into value
			priority := getPriority(item)
			
			// Set bitmask
			cpmOneBitmask |= uint64(1 << priority)
		}

		for _, item := range compartmentTwo {
			// Convert each item into value
			priority := getPriority(item)
			
			// Set bitmask
			cpmTwoBitmask |= uint64(1 << priority)
		}

		// AND bitmask
		duplicateItems := cpmOneBitmask & cpmTwoBitmask

		// fmt.Printf("%64b\n", cpmOneBitmask)
		// fmt.Printf("%64b\n", cpmTwoBitmask)
		// fmt.Printf("%64b\n", duplicateItems)

		// Get result
		for key := 1; key <= 52; key++ {
			if duplicateItems & uint64(1 << key) != 0 {
				total += key
			}
		}
	}

	fmt.Printf("Part 1: Total priorities = %d\n", total)

	fmt.Printf("Part 2: Total priorities = %d\n", resultP2)
}

func calculatePart2(scanner *bufio.Scanner) int {
	var jointBitmask uint64
	

	
	result := 0

	for scanner.Scan() {
		for n := 0; n <= 63; n++ {
			jointBitmask |= 1 << n
		}
		fmt.Printf("Start %64b\n", jointBitmask)


		// Get three lines
		for n := 1; n <= 3; n++ {
			var lineBitmask uint64

			if n != 1 {
				scanner.Scan()
			}
			line := scanner.Text()

			for _, item := range line {
				// Convert each item into value
				priority := getPriority(item)
				
				// Set bitmask
				lineBitmask |= uint64(1 << priority)
			}

			jointBitmask &= lineBitmask
			
			fmt.Printf(" Pass %64b\n", lineBitmask)
			fmt.Printf(" Res  %64b\n", jointBitmask)
		}


		for key := 1; key <= 52; key++ {
			if jointBitmask & uint64(1 << key) != 0 {
				result += key
			}
		}
	}


	return result
}

func getPriority(item rune) int {
	priority := int(item)

	if priority < (26 + 65) {
		// Uppercase
		// A = 65, should be 27
		priority -= (65 - 27)
	} else {
		// Lowercase
		// a = 97, should be 1
		priority -= 96
	}

	return priority
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