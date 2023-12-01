package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type digit_at_fn func(string, int) int

// get the calibration value, error => -1
func getCalibrationValue(line string, digit_at digit_at_fn) int {
	// search first digit from the front
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

	// search last digit from the back
	for i := len(line) - 1; i >= 0; i-- {
		last := digit_at(line, i)
		if last >= 0 {
			return first*10 + last
		}
	}

	return -1
}

// check if s[i] is a digit;
// return 0-9 if digit is found, otherwise -1
func DigitAt(s string, pos int) int {
	c := s[pos]
	if '0' <= c && c <= '9' {
		return int(c - '0')
	}
	return -1
}

// check if sub can be found in main at pos
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

var one_to_nine = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// check if s[i] is a digit, or a "digit name" starts at s[i]
// return 0-9 if a digt is found, -1 otherwise
func DigitOrNameAt(s string, pos int) int {
	// look for digits first
	c := DigitAt(s, pos)
	if c >= 0 {
		return c
	}

	// check all "digit names"
	for i, name := range one_to_nine {
		if IsSubstringAt(name, s, pos) {
			return i + 1
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
			c = getCalibrationValue(line, DigitOrNameAt)
		}
		if c < 0 {
			panic(fmt.Errorf("bad line: %s", line))
		}
		calibrationValueSum += c
	})

	fmt.Println(calibrationValueSum)
}
