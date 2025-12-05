package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	low, high int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var freshAvailable, freshAll int
	ranges := make([]Range, 0)
	rangesLoaded := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			rangesLoaded = true
			slices.SortFunc(ranges, func(a, b Range) int {
				return a.low - b.low
			})
			continue
		} else if !rangesLoaded {
			r := strings.Split(line, "-")
			low, _ := strconv.Atoi(r[0])
			high, _ := strconv.Atoi(r[1])
			ranges = append(ranges, Range{low: low, high: high})
		} else {
			v, _ := strconv.Atoi(line)
			for i := range ranges {
				if ranges[i].low <= v && v <= ranges[i].high {
					freshAvailable++
					break
				}
			}
		}
	}

	var prevHigh int
	for i := range ranges {
		r := ranges[i]
		if r.high < prevHigh {
			continue
		}
		freshAll += (r.high - max(r.low, prevHigh+1)) + 1
		prevHigh = r.high
	}

	fmt.Printf("Fresh Available: %d\n", freshAvailable)
	fmt.Printf("Fresh All: %d\n", freshAll)
}
