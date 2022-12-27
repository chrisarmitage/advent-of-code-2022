package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)


func main() {
	elves := make(map[int]int)
	readFile, err := os.Open("input.txt")
  
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

	fmt.Printf("Single elf with most calories: Elf %d with %d calories\n", keys[0], elves[keys[0]])
    
	topCalories := 0
	for i := 0; i < 3; i++ {
		topCalories += elves[keys[i]]
	}

	fmt.Printf("Calories carried by top three elves: %d calories", topCalories)
}