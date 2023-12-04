package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2023/helpers"
)

var re_spaces = regexp.MustCompile(` +`)

func cardMatches(card string) int {
	//         /   winning  \   /       numbers       \
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	matches := 0
	winning := make(map[string]bool)
	pipe := false
	for i, token := range re_spaces.Split(card, -1) {
		if i >= 2 {
			if token == "|" {
				pipe = true
			} else if !pipe {
				winning[token] = true
			} else {
				if _, ok := winning[token]; ok {
					matches += 1
				}
			}
		}
	}
	return matches
}

func cardScore(card string) int {
	m := cardMatches(card)
	if m == 0 {
		return 0
	}
	return 1 << (m - 1)
}

func initIfNeeded(count []int, index int) []int {
	for index >= len(count) {
		count = append(count, 1)
	}
	return count
}

func main() {
	sum := 0
	if helpers.Part1() {
		helpers.ReadStdin(func(line string) {
			sum += cardScore(line)
		})
	} else {
		count := make([]int, 0)
		index := 0
		helpers.ReadStdin(func(line string) {
			count = initIfNeeded(count, index)
			sum += count[index]
			matches := cardMatches(line)
			for i := 1; i <= matches; i += 1 {
				count = initIfNeeded(count, index+i)
				count[index+i] += count[index]
			}
			index += 1
		})
	}
	fmt.Println(sum)
}
