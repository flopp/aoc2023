package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

func hash(s string) int {
	h := 0
	for _, r := range s {
		h += int(r)
		h *= 17
		h %= 256
	}
	return h
}

type Box struct {
	Order  []string
	Values map[string]int
}

func main() {
	steps := ([]string)(nil)
	helpers.ReadStdin(func(line string) {
		steps = strings.Split(line, ",")
	})

	if helpers.Part1() {
		sum := 0
		for _, step := range steps {
			sum += hash(step)
		}
		fmt.Println(sum)
	} else {
		boxes := make([]*Box, 256)
		for i := 0; i < 256; i += 1 {
			boxes[i] = &Box{make([]string, 0), make(map[string]int)}
		}
		for _, step := range steps {
			if step[len(step)-1] == '-' {
				name := step[0 : len(step)-1]
				b := boxes[hash(name)]
				if _, ok := b.Values[name]; ok {
					delete(b.Values, name)
					for i, n := range b.Order {
						if n == name {
							for j := i + 1; j < len(b.Order); j += 1 {
								b.Order[j-1] = b.Order[j]
							}
							b.Order = b.Order[:len(b.Order)-1]
						}
					}
				}
			} else {
				name := step[0 : len(step)-2]
				b := boxes[hash(name)]
				value := helpers.MustParseInt(step[len(step)-1:])
				if _, ok := b.Values[name]; !ok {
					b.Order = append(b.Order, name)
				}
				b.Values[name] = value
			}
		}

		power := 0
		for boxi, box := range boxes {
			for i, name := range box.Order {
				power += (1 + boxi) * (1 + i) * box.Values[name]
			}
		}
		fmt.Println(power)
	}
}
