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

	score := 0

	var line string
	var movePoints int
	var resultPoints int

	// A X = R
	// B Y = P
	// C Z = S
    for fileScanner.Scan() {
        line = fileScanner.Text()

		s := strings.Split(line, " ")
    	theirMove, myMove := s[0], s[1]

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

		score += (movePoints + resultPoints)
	}
	
    readFile.Close()

	fmt.Printf("Final score: %d\n", score)
    

}