package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

type NodeInfo struct {
	left, right string
}
type Node struct {
	name        string
	left, right *Node
	start       bool
	target      bool
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func main() {
	re_directions := regexp.MustCompile(`^(L|R)+$`)
	re_node := regexp.MustCompile(`^(.+) = \((.+), (.+)\)$`)
	directions := ""

	nodeInfos := make(map[string]*NodeInfo)
	helpers.ReadStdin(func(line string) {
		if m := re_directions.FindStringSubmatch(line); m != nil {
			directions = line
		} else if m := re_node.FindStringSubmatch(line); m != nil {
			nodeInfos[m[1]] = &NodeInfo{m[2], m[3]}
		}
	})

	transitions := make(map[string]*Node)
	for name, _ := range nodeInfos {
		if helpers.Part1() {
			transitions[name] = &Node{name, nil, nil, name == "AAA", name == "ZZZ"}
		} else {
			transitions[name] = &Node{name, nil, nil, strings.HasSuffix(name, "A"), strings.HasSuffix(name, "Z")}
		}
	}
	for name, info := range nodeInfos {
		transitions[name].left = transitions[info.left]
		transitions[name].right = transitions[info.right]
	}

	if helpers.Part1() {
		node := transitions["AAA"]
		for step := 0; true; step += 1 {
			if node.target {
				fmt.Println(step)
				break
			}

			if directions[step%len(directions)] == 'L' {
				node = node.left
			} else {
				node = node.right
			}
		}
	} else {
		nodes := make([]*Node, 0)
		for _, node := range transitions {
			if node.start {
				nodes = append(nodes, node)
			}
		}

		var commonSteps int64 = 1

		for _, node := range nodes {
			for step := 0; true; step += 1 {
				if node.target {
					commonSteps = lcm(commonSteps, int64(step))
					break
				}
				if directions[step%len(directions)] == 'L' {
					node = node.left
				} else {
					node = node.right
				}
			}
		}

		fmt.Println(commonSteps)
	}
}
