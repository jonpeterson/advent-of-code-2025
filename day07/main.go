package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Splitter struct {
	exists bool
	paths  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	startIndex := -1
	splitters := make([][]Splitter, 0)

	scanner := bufio.NewScanner(file)
	var x, y int
	for scanner.Scan() {
		line := scanner.Text()
		if startIndex == -1 {
			startIndex = strings.IndexRune(line, 'S')
		} else if strings.ContainsRune(line, '^') {
			x = 0
			splitters = append(splitters, make([]Splitter, y+1))
			for i := startIndex - y; i <= startIndex+y; i += 2 {
				if line[i] == '^' {
					(&splitters[y][x]).exists = true
				}
				x++
			}
			y++
		}
	}

	splitterPaths := evalBeamPosition(splitters, 0, 0)

	var splittersUsed int
	for y = range splitters {
		for x = range splitters[y] {
			if splitters[y][x].paths > 0 {
				splittersUsed++
			}
		}
	}

	fmt.Printf("Splits: %d\n", splittersUsed)
	fmt.Printf("Paths: %d\n", splitterPaths)
}

func evalBeamPosition(splitters [][]Splitter, x int, y int) int {
	if y >= len(splitters) {
		return 1
	}
	splitter := &splitters[y][x]
	if !splitter.exists {
		return evalBeamPosition(splitters, x+1, y+2)
	} else if splitter.paths == 0 {
		a := evalBeamPosition(splitters, x, y+1)
		b := evalBeamPosition(splitters, x+1, y+1)
		splitter.paths = a + b
	}
	return splitter.paths
}
