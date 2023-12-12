package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

func arrangements(springs []byte, springIndex int, defects []int, defectIndex, defectRem int) int {
	if springIndex == len(springs) {
		if defectIndex == len(defects) {
			return 1
		}
		if defectIndex+1 == len(defects) && defectRem == 0 {
			return 1
		}
		return 0
	}
	switch springs[springIndex] {
	case '#':
		if defectRem < 0 {
			if defectIndex == len(defects) {
				return 0
			}
			defectRem = defects[defectIndex]
		}
		defectRem -= 1
		if defectRem < 0 {
			return 0
		}
		return arrangements(springs, springIndex+1, defects, defectIndex, defectRem)
	case '.':
		if defectRem > 0 {
			return 0
		}
		if defectRem == 0 {
			defectRem = -1
			defectIndex += 1
		}
		return arrangements(springs, springIndex+1, defects, defectIndex, defectRem)
	case '?':
		sum := 0
		// .
		springs[springIndex] = '.'
		sum += arrangements(springs, springIndex, defects, defectIndex, defectRem)
		// #
		springs[springIndex] = '#'
		sum += arrangements(springs, springIndex, defects, defectIndex, defectRem)
		springs[springIndex] = '?'

		return sum
	}
	return 0
}

type Line struct {
	springs []byte
	defects []int
}

func main() {
	lines := make([]Line, 0)
	helpers.ReadStdin(func(line string) {
		a := strings.Split(line, " ")
		springs := make([]byte, 0)
		for i := 0; i < len(a[0]); i += 1 {
			springs = append(springs, a[0][i])
		}
		defects := make([]int, 0)
		for _, token := range strings.Split(a[1], ",") {
			defects = append(defects, helpers.MustParseInt(token))
		}
		if !helpers.Part1() {
			// unfold
			springs5 := make([]byte, 0, 6*len(springs))
			for unfold := 0; unfold < 5; unfold += 1 {
				if unfold > 0 {
					springs5 = append(springs5, '?')
				}
				springs5 = append(springs5, springs...)
			}
			defects5 := make([]int, 0, 5*len(defects))
			for unfold := 0; unfold < 5; unfold += 1 {
				defects5 = append(defects5, defects...)
			}

			springs = springs5
			defects = defects5
		}
		lines = append(lines, Line{springs, defects})
	})

	sum := 0
	for _, line := range lines {
		defectsSum := 0
		for _, d := range line.defects {
			defectsSum += d
		}
		s := arrangements(line.springs, 0, line.defects, 0, -1)
		sum += s
	}
	fmt.Println(sum)
}
