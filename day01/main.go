package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type digit_at_fn func(string, int) int

// error => -1
func getCalibrationValue(line string, digit_at digit_at_fn) int {
	first := -1
	for i := 0; i < len(line); i++ {
		first = digit_at(line, i)
		if first >= 0 {
			break
		}
	}
	if first < 0 {
		return -1
	}

	for i := len(line) - 1; i >= 0; i-- {
		last := digit_at(line, i)
		if last >= 0 {
			return first*10 + last
		}
	}

	return -1
}

func DigitAt(s string, pos int) int {
	c := s[pos]
	if '0' <= c && c <= '9' {
		return int(c - '0')
	}
	return -1
}

func IsSubstringAt(sub, main string, pos int) bool {
	sublen := len(sub)
	mainlen := len(main)
	if sublen+pos > mainlen {
		return false
	}
	for i := 0; i < sublen; i += 1 {
		if main[pos+i] != sub[i] {
			return false
		}
	}
	return true
}

var digits = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// error => -1
func DigitNameAt(s string, pos int) int {
	c := DigitAt(s, pos)
	if c >= 0 {
		return c
	}

	for j, d := range digits {
		if j != 0 && IsSubstringAt(d, s, pos) {
			return j
		}
	}

	return -1
}

func main() {
	calibrationValueSum := 0
	helpers.ReadStdin(func(line string) {
		c := 0
		if helpers.Part1() {
			c = getCalibrationValue(line, DigitAt)
		} else {
			c = getCalibrationValue(line, DigitNameAt)
		}
		if c < 0 {
			panic(fmt.Errorf("bad line: %s", line))
		}
		calibrationValueSum += c
	})

	fmt.Println(calibrationValueSum)

}
