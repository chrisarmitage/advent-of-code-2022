package main

import (
	"bufio"
	"fmt"
	"os"
)

const SAMPLE_TARGET_P1 = "CMZ"

const STACKS = 9

type Move struct {
	Qty  int
	From int
	To   int
}

func main() {
	scanner := makeScanner("input.txt")
	resultP1 := calculatePart1(scanner)

	fmt.Printf("Part 1: Top containers = %s\n", resultP1)

}

func calculatePart1(scanner *bufio.Scanner) string {
	stacks := make([][]string, STACKS)

	for col := 0; col < STACKS; col++ {
		stacks[col] = []string{}
	}

	mode := "setup"
	for scanner.Scan() {
		line := scanner.Text()

		if mode == "setup" {
			if string(line[1]) == "1" {
				mode = "skip"
				continue
			}
			for col := 0; col < STACKS; col++ {
				crateLetter := string(line[(col*4)+1])
				fmt.Printf("%s ", crateLetter)

				if crateLetter != " " {
					stacks[col] = append([]string{crateLetter}, stacks[col]...)
				}
			}
			fmt.Println()
		}

		if mode == "skip" {
			if line != "" {
				mode = "move"
				fmt.Printf("Init : %v\n", stacks)
			}
		}

		if mode == "move" {
			fmt.Printf("Action: %s\n", line)

			var m Move

			fmt.Sscanf(
				line, "move %d from %d to %d",
				&m.Qty, &m.From, &m.To)

			m.From = m.From - 1
			m.To = m.To - 1

			for n := 1; n <= m.Qty; n++ {
				fromStackSize := len(stacks[m.From]) - 1
				crateToMove := stacks[m.From][fromStackSize]

				// Pop
				stacks[m.From] = append(stacks[m.From][:fromStackSize], stacks[m.From][fromStackSize+1:]...)

				// Push
				stacks[m.To] = append(stacks[m.To], crateToMove)

				fmt.Printf("Stacks: %v\n", stacks)
			}
		}
	}

	fmt.Printf("Final Stacks: %v\n", stacks)

	result := ""
	for col := 0; col < STACKS; col++ {
		fromStackSize := len(stacks[col]) - 1
		topCrate := stacks[col][fromStackSize]
		result += topCrate
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
