package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total2 := 0
	total12 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total2 += greatestValue(line, 2)
		total12 += greatestValue(line, 12)
	}

	fmt.Printf("Total 2: %d\n", total2)
	fmt.Printf("Total 12: %d\n", total12)
}

func greatestValue(s string, c int) int {
	var si, sv, cv int

	for i := c - 1; i >= 0; i-- {
		si = greatestRuneIndex(s, si, len(s)-i)
		cv, _ = strconv.Atoi(s[si : si+1])
		sv = (sv * 10) + cv
		si++
	}

	return sv
}

func greatestRuneIndex(s string, start int, end int) int {
	var gri int
	var r, gr rune

	for i := start; i < end; i++ {
		r = rune(s[i])
		if r > gr {
			gr = r
			gri = i
		}
	}

	return gri
}
