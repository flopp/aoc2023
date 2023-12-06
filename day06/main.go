package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

func waysToBeat(time, recordDistance int) int {
	count := 0
	for holdTime := 0; holdTime <= time; holdTime += 1 {
		if holdTime*(time-holdTime) > recordDistance {
			count += 1
		}
	}
	return count
}

var re_spaces = regexp.MustCompile(` +`)

func main() {
	if helpers.Part1() {
		times := make([]int, 0)
		recordDistances := make([]int, 0)
		helpers.ReadStdin(func(line string) {
			values := make([]int, 0)
			for _, token := range re_spaces.Split(line, -1)[1:] {
				values = append(values, helpers.MustParseInt(token))
			}
			if len(times) == 0 {
				times = values
			} else {
				recordDistances = values
			}
		})

		product := 1
		for i, time := range times {
			product *= waysToBeat(time, recordDistances[i])
		}
		fmt.Println(product)
	} else {
		time := 0
		recordDistance := 0
		helpers.ReadStdin(func(line string) {
			value := helpers.MustParseInt(strings.Join(re_spaces.Split(line, -1)[1:], ""))
			if time == 0 {
				time = value
			} else {
				recordDistance = value
			}
		})

		fmt.Println(waysToBeat(time, recordDistance))
	}
}
