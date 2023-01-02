package day07

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Name       string
	ActualSize int
	Children   []*Entry
}

//var root *Entry

func newEntry(name string) *Entry {
	return &Entry{
		Name:     name,
		//Children: []Entry,
	}
}

func (e *Entry) AddChild(child *Entry) {
	e.Children = append(e.Children, child)
}

func (e *Entry) getSize() int {
	if len(e.Children) > 0 {
		size := 0
		for _, child := range e.Children {
			size += child.getSize()
		}

		return size
	}
	return e.ActualSize
} 

var printOffset = 1

func (e *Entry) Print() {
	fmt.Printf("%s%s [%d]\n", strings.Repeat("- ", printOffset), e.Name, e.getSize())
	if len(e.Children) > 0 {
		printOffset += 1
		// fmt.Println("Increased")
		for _, child := range e.Children {
			child.Print()
		}
		printOffset -= 1
		// fmt.Println("Decreased")
	}
} 

var totalP1 = 0

func (e *Entry) CalculateTotalSize(upperLimit int) int {
	totalSize := 0
	if len(e.Children) > 0 {
		for _, child := range e.Children {
			totalSize += child.CalculateTotalSize(upperLimit)
		}

		if totalSize <= upperLimit {
			fmt.Printf("Directory %s has a size of %d\n", e.Name, totalSize)
			totalP1 += totalSize
		}

		return totalSize

	}

	return e.ActualSize

} 

func Run(sample bool) (string, error) {
	filename := "day07/input.txt"
	if sample {
		filename = "day07/sample.txt"
	}

	scanner, err := makeScanner(filename)
	if err != nil {
		return "", err
	}

	root := parseInput(scanner)

	root.CalculateTotalSize(100000)

	//root := newEntry("/")

	// a := newEntry("a")
	// root.AddChild(a)
	
	// b := newEntry("b.txt")
	// b.ActualSize = 14848514
	// root.AddChild(b)
	
	// c := newEntry("c.dat")
	// c.ActualSize = 8504156
	// root.AddChild(c)
	
	// d := newEntry("d")
	// root.AddChild(d)

	// fmt.Printf("%+v\n", root)

	// fmt.Printf("Size /: %d\n", root.getSize())
	// fmt.Printf("Size A: %d\n", a.getSize())
	// fmt.Printf("Size B: %d\n", b.getSize())
	// fmt.Printf("Size C: %d\n", c.getSize())
	// fmt.Printf("Size D: %d\n", d.getSize())

	root.Print()

	return fmt.Sprintf(
			"Part 1: Total filesize = %d\n", 
			totalP1,
		),
		 nil
}

func makeScanner(filename string) (*bufio.Scanner, error) {
	readFile, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func parseInput(scanner *bufio.Scanner) *Entry {
	// state := "ready"
	dirStack := make([]*Entry, 0)
	
	root := newEntry("/")
	currentDirectory := root
	dirStack = append(dirStack, currentDirectory)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0:1] == "$" {
			// state = "command"
			// Command
			if line[2:4] == "cd" {
				directoryName := line[5:]
				if directoryName == ".." {
					dirStack = dirStack[:len(dirStack)-1]
					currentDirectory = dirStack[len(dirStack)-1:][0]
					fmt.Printf("Debug: Moving up to %s\n", currentDirectory.Name)
				} else if directoryName != "/" {
					newDirectory := newEntry(directoryName)
					fmt.Printf("Debug: Moving to %s\n", newDirectory.Name)
					currentDirectory.AddChild(newDirectory)
					fmt.Printf("Debug: Added %s to %s\n", newDirectory.Name, currentDirectory.Name)

					dirStack = append(dirStack, newDirectory)
					currentDirectory = newDirectory
				}
				
				// fmt.Println("Dumping directory tree")
				// root.Print()
			} else if line[2:4] == "ls" {
				// state = "listing"
			}

		} else {
			// Listing
			var fileSize int
			var filename string

			fmt.Sscanf(
				line, "%d %s",
				&fileSize,
				&filename)
			fmt.Printf("[line, size, name] [%s, %d, %s]\n", line, fileSize, filename)

			if filename != "" {
				newFile := newEntry(filename)
				newFile.ActualSize = fileSize
				fmt.Printf("Debug: Adding %s to %s\n", newFile.Name, currentDirectory.Name)
				currentDirectory.AddChild(newFile)
			}
		}
	}

	return root
}