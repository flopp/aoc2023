package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

func isDigit(line string, x int) bool {
	c := line[x]
	return '0' <= c && c <= '9'
}

func isSymbol(line string, x int) bool {
	c := line[x]
	return c != '.' && (c < '0' || c > '9')
}

func hasAdjacentSymbol(lines []string, minx, maxx, y int) bool {
	for yy := helpers.Max(y-1, 0); yy <= helpers.Min(y+1, len(lines)-1); yy += 1 {
		line := lines[yy]
		for xx := helpers.Max(minx-1, 0); xx <= helpers.Min(maxx+1, len(line)-1); xx += 1 {
			if isSymbol(line, xx) {
				return true
			}
		}
	}

	return false
}

func getNumber(line string, x int) (int, int, int) {
	if !isDigit(line, x) {
		return -1, 0, 0
	}

	// find start and end of number
	minx := x
	for minx > 0 && isDigit(line, minx-1) {
		minx -= 1
	}
	maxx := x
	for maxx+1 < len(line) && isDigit(line, maxx+1) {
		maxx += 1
	}
	return helpers.MustParseInt(line[minx : 1+maxx]), minx, maxx
}

func getAdjacentNumbers(lines []string, x, y int) []int {
	nums := make([]int, 0)
	for yy := helpers.Max(0, y-1); yy <= helpers.Min(len(lines)-1, y+1); yy += 1 {
		line := lines[yy]
		n, _, _ := getNumber(line, x)
		if n != -1 {
			nums = append(nums, n)
		} else {
			for xx := helpers.Max(0, x-1); xx <= helpers.Min(len(line)-1, x+1); xx += 1 {
				n, _, _ = getNumber(line, xx)
				if n != -1 {
					nums = append(nums, n)
				}
			}
		}
	}
	return nums
}

func main() {
	lines := make([]string, 0)
	helpers.ReadStdin(func(line string) {
		lines = append(lines, line)
	})
	w := len(lines[0])

	sum := 0
	for y, line := range lines {
		if helpers.Part1() {
			for x := 0; x < w; x += 1 {
				n, minx, maxx := getNumber(line, x)
				if n >= 0 {
					if hasAdjacentSymbol(lines, minx, maxx, y) {
						sum += n
					}
					x = maxx
				}
			}
		} else {
			for x := 0; x < w; x += 1 {
				if line[x] == '*' {
					nums := getAdjacentNumbers(lines, x, y)
					if len(nums) == 2 {
						sum += nums[0] * nums[1]
					}
				}
			}
		}
	}
	fmt.Println(sum)
}
