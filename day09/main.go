package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

func computeDiffs(numbers *[]int) bool {
	allZero := true
	newlen := len(*numbers) - 1
	for i := 0; i < newlen; i += 1 {
		d := (*numbers)[i+1] - (*numbers)[i]
		if d != 0 {
			allZero = false
		}
		(*numbers)[i] = d
	}
	(*numbers) = (*numbers)[:newlen]
	return allZero
}

func nextNumber(numbers []int) int {
	lastsum := 0
	diffs := make([]int, 0, len(numbers))
	diffs = append(diffs, numbers...)
	allZero := false
	for !allZero {
		lastsum += diffs[len(diffs)-1]
		allZero = computeDiffs(&diffs)
	}
	return lastsum
}

func prevNumber(numbers []int) int {
	first := make([]int, 0, len(numbers))
	diffs := make([]int, 0, len(numbers))
	diffs = append(diffs, numbers...)
	allZero := false
	for !allZero {
		first = append(first, diffs[0])
		allZero = computeDiffs(&diffs)
	}

	n := 0
	for i := len(first) - 1; i >= 0; i -= 1 {
		n = first[i] - n
	}
	return n
}

func main() {
	sum := 0
	helpers.ReadStdin(func(line string) {
		numbers := make([]int, 0)
		for _, token := range strings.Split(line, " ") {
			numbers = append(numbers, helpers.MustParseInt(token))
		}
		if helpers.Part1() {
			sum += nextNumber(numbers)
		} else {
			sum += prevNumber(numbers)
		}
	})
	fmt.Println(sum)
}
