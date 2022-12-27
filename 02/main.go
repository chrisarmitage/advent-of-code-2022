package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	scoreP1 := 0
	scoreP2 := 0

	var line string
	var movePoints int
	var resultPoints int
	var p2Points int

	// A X = R
	// B Y = P
	// C Z = S

	// Required result
	// X = Lose
	// Y = Draw
	// Z = Win
    for fileScanner.Scan() {
        line = fileScanner.Text()

		s := strings.Split(line, " ")
    	theirMove, myMove := s[0], s[1]
		targetResult := s[1]

		if (
			(theirMove == "A" && myMove == "X") ||
			(theirMove == "B" && myMove == "Y") ||
			(theirMove == "C" && myMove == "Z")) {
				// Draw
				resultPoints = 3
		} else if (
			(theirMove == "A" && myMove == "Y") ||
			(theirMove == "B" && myMove == "Z") ||
			(theirMove == "C" && myMove == "X")) {
				// Win
				resultPoints = 6
		} else {
			resultPoints = 0
		}

		if (myMove == "X") {
			movePoints = 1
		} else if (myMove == "Y") {
			movePoints = 2
		} else {
			movePoints = 3
		}

		// Strategy 2
		// A X = R
		// B Y = P
		// C Z = S

		// Required result
		// X = Lose
		// Y = Draw
		// Z = Win
		if (theirMove == "A") {
			if (targetResult == "X") {
				p2Points = 3 + 0
			} else if (targetResult == "Y") {
				p2Points = 1 + 3
			} else {
				p2Points = 2 + 6
			}
		} else if (theirMove == "B") {
			if (targetResult == "X") {
				p2Points = 1 + 0
			} else if (targetResult == "Y") {
				p2Points = 2 + 3
			} else {
				p2Points = 3 + 6
			}
		} else {
			if (targetResult == "X") {
				p2Points = 2 + 0
			} else if (targetResult == "Y") {
				p2Points = 3 + 3
			} else {
				p2Points = 1 + 6
			}
		}


		scoreP1 += (movePoints + resultPoints)
		scoreP2 += p2Points
	}
	
    readFile.Close()

	fmt.Printf("Final score (strategy one): %d\n", scoreP1)
    fmt.Printf("Final score (strategy two): %d\n", scoreP2)
    

}