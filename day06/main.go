package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	lineOffset int
	operator   rune
	lrValues   []int
	tdValues   []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	problems := make([]Problem, 0)

	scanner := bufio.NewScanner(file)
	valueLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		fr := line[0]
		if fr == '+' || fr == '*' {
			for i, r := range line {
				if r == '+' || r == '*' {
					problems = append(problems, Problem{lineOffset: i, operator: r, lrValues: make([]int, 0), tdValues: make([]int, 0)})
				}
			}
			break
		} else {
			valueLines = append(valueLines, line)
		}
	}

	for pi, p := range problems {
		lo := p.lineOffset

		for _, vl := range valueLines {
			vs := ""
			for li := lo; ; li++ {
				r := ' '
				if li < len(vl) {
					r = int32(vl[li])
				}
				if r != ' ' {
					vs += string(r)
				} else if vs != "" {
					break
				}
			}
			v, _ := strconv.Atoi(vs)
			p.lrValues = append(p.lrValues, v)
		}

		for li := lo; ; li++ {
			vs := ""
			for _, vl := range valueLines {
				r := ' '
				if li < len(vl) {
					r = int32(vl[li])
				}
				if r != ' ' {
					vs += string(r)
				}
			}
			if vs == "" {
				break
			}
			v, _ := strconv.Atoi(vs)
			p.tdValues = append(p.tdValues, v)
		}

		problems[pi] = p
	}

	var lrTotal int
	var tdTotal int
	for _, p := range problems {
		var lra int
		var tda int
		switch p.operator {
		case '+':
			lra = 0
			tda = 0
			for _, v := range p.lrValues {
				lra += v
			}
			for _, v := range p.tdValues {
				tda += v
			}
		case '*':
			lra = 1
			tda = 1
			for _, v := range p.lrValues {
				lra *= v
			}
			for _, v := range p.tdValues {
				tda *= v
			}
		}
		lrTotal += lra
		tdTotal += tda
	}

	fmt.Printf("LR Total: %d\n", lrTotal)
	fmt.Printf("TD Total: %d\n", tdTotal)
}
