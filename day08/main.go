package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2023/helpers"
)

type Node struct {
	left, right string
}

func main() {
	if !helpers.Part1() {
		panic("part2 is not implemented, yet")
	}

	re_directions := regexp.MustCompile(`^(L|R)+$`)
	re_node := regexp.MustCompile(`^(.+) = \((.+), (.+)\)$`)
	directions := ""
	nodes := make(map[string]*Node)

	helpers.ReadStdin(func(line string) {
		if m := re_directions.FindStringSubmatch(line); m != nil {
			directions = line
		} else if m := re_node.FindStringSubmatch(line); m != nil {
			nodes[m[1]] = &Node{m[2], m[3]}
		} else if line != "" {
			panic(fmt.Errorf("bad line: <%s>", line))
		}
	})

	node := "AAA"
	for step := 0; true; step += 1 {
		if node == "ZZZ" {
			fmt.Println(step)
			break
		}

		if directions[step%len(directions)] == 'L' {
			node = nodes[node].left
		} else {
			node = nodes[node].right
		}
	}
}
