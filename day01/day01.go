package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Run(sample bool) (string, error) {
	elves := make(map[int]int)
	filename := "day01/input.txt"
	if sample {
		filename = "day01/sample.txt"
	}
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var line string
	elfId := 1
	elfCalories := 0
	for fileScanner.Scan() {
		line = fileScanner.Text()

		if line == "" {
			elves[elfId] = elfCalories
			elfCalories = 0
			elfId++
		} else {
			i, _ := strconv.Atoi(line)
			elfCalories += i
		}
	}

	readFile.Close()

	keys := make([]int, 0, len(elves))
	for key := range elves {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return elves[keys[i]] > elves[keys[j]] })

	topCalories := 0
	for i := 0; i < 3; i++ {
		topCalories += elves[keys[i]]
	}

	return fmt.Sprintf(
		"Single elf with most calories: Elf %d with %d calories\n" +
		"Calories carried by top three elves: %d calories",
		 keys[0],
		  elves[keys[0]],
		   topCalories,
		),
		nil

}