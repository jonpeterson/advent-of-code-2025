package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	twiceTotal := 0
	atLeastTwiceTotal := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(splitComma)
	for scanner.Scan() {
		r := strings.Split(scanner.Text(), "-")
		low, _ := strconv.Atoi(r[0])
		high, _ := strconv.Atoi(r[1])

		for _, i := range findDoubles(low, high) {
			twiceTotal += i
		}

		for _, i := range findRepeats(low, high) {
			fmt.Printf("Adding %d\n", i)
			atLeastTwiceTotal += i
		}
	}

	fmt.Printf("Twice Total: %d\n", twiceTotal)
	fmt.Printf("At Least Twice Total: %d\n", atLeastTwiceTotal)
}

func findDoubles(low int, high int) []int {
	res := make([]int, 0)

	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		sl := len(s)
		if (sl % 2) != 0 {
			continue
		}
		hsl := sl / 2
		if s[:hsl] == s[hsl:] {
			res = append(res, i)
		}
	}

	return res
}

func findRepeats(low int, high int) []int {
	res := make([]int, 0)

	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		sl := len(s)
		hsl := sl / 2
	j:
		for j := 1; j <= hsl; j++ {
			if (sl % j) != 0 {
				continue
			}
			pat := s[:j]
			for k := j; k < sl; k += j {
				if s[k:k+j] != pat {
					continue j
				}
			}
			res = append(res, i)
			break
		}
	}

	return res
}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
