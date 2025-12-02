package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dial struct {
	Current     int
	PassedZero  int
	EndedOnZero int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dial := Dial{Current: 50}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := rune(line[0])
		amt, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if dir == 'L' {
			amt *= -1
		}

		prev := dial.Current

		dial.Turn(amt)

		fmt.Printf("%s: %d --[%d]-> %d (zeros: %d)\n", line, prev, amt, dial.Current, dial.PassedZero)
	}

	fmt.Printf("End on zero count: %d\n", dial.EndedOnZero)
	fmt.Printf("Pass zero count: %d\n", dial.PassedZero)
	fmt.Printf("Dial ends on: %d\n", dial.Current)
}

func (dial *Dial) Turn(amt int) {
	next := dial.Current + amt

	if next < 1 {
		if dial.Current != 0 {
			dial.PassedZero++
		}
		dial.PassedZero += (amt + dial.Current) / -100
	} else if next > 99 {
		dial.PassedZero += (amt + dial.Current) / 100
	}

	dial.Current = next % 100
	if dial.Current < 0 {
		dial.Current += 100
	}

	if dial.Current == 0 {
		dial.EndedOnZero++
	}
}
