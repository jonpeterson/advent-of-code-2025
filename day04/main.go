package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]bool, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))

		for i, ch := range line {
			if ch == '@' {
				row[i] = true
			}
		}

		grid = append(grid, row)
	}

	var c, grandTotal int
	var r []bool
	for pass := 1; ; pass++ {
		newGrid := make([][]bool, len(grid))
		var total int
		for i := range grid {
			newRow := slices.Clone(grid[i])
			newGrid[i] = newRow
			for j := range grid[i] {
				if !grid[i][j] {
					continue
				}

				c = 0

				if i > 0 {
					r = grid[i-1]
					if j > 0 && r[j-1] {
						c++
					}
					if r[j] {
						c++
					}
					if j < len(r)-1 && r[j+1] {
						c++
					}
				}

				r = grid[i]
				if j > 0 && r[j-1] {
					c++
				}
				if j < len(r)-1 && r[j+1] {
					c++
				}

				if i < len(grid)-1 {
					r = grid[i+1]
					if j > 0 && r[j-1] {
						c++
					}
					if r[j] {
						c++
					}
					if j < len(r)-1 && r[j+1] {
						c++
					}
				}

				if c < 4 {
					total++
					newRow[j] = false
				}
			}
		}

		if total == 0 {
			break
		}

		fmt.Printf("Pass %d Total: %d\n", pass, total)
		grandTotal += total

		grid = newGrid
	}

	fmt.Printf("Grand Total: %d\n", grandTotal)
}
