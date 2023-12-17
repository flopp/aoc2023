package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type Dir int

const (
	N = Dir(0)
	E = Dir(1)
	S = Dir(2)
	W = Dir(3)
)

func (d Dir) rot(offset int) Dir {
	r := int(d) + offset
	for r < 0 {
		r += 4
	}
	return Dir(r % 4)
}

type XY struct {
	x, y int
}

func (xy XY) adv(d Dir) XY {
	switch d {
	case N:
		return XY{xy.x, xy.y - 1}
	case E:
		return XY{xy.x + 1, xy.y}
	case S:
		return XY{xy.x, xy.y + 1}
	case W:
		return XY{xy.x - 1, xy.y}
	}
	return xy
}

type Grid struct {
	data []int
	w, h int
}

func (g *Grid) add(row string) {
	g.w = len(row)
	g.h += 1

	for _, c := range row {
		g.data = append(g.data, helpers.MustParseInt(string(c)))
	}
}

func (g Grid) valid(xy XY) bool {
	return xy.x >= 0 && xy.y >= 0 && xy.x < g.w && xy.y < g.h
}

func (g *Grid) next(n *Node, mindlen int, maxdlen int) []*Node {
	if n.dlen < mindlen {
		res := make([]*Node, 0, 1)
		if xy := n.xy.adv(n.d); g.valid(xy) {
			res = append(res, &Node{xy, n.d, n.dlen + 1, n.heatLoss + g.data[xy.y*g.w+xy.x]})
		}
		return res
	}

	res := make([]*Node, 0, 3)
	if n.dlen < maxdlen {
		if xy := n.xy.adv(n.d); g.valid(xy) {
			res = append(res, &Node{xy, n.d, n.dlen + 1, n.heatLoss + g.data[xy.y*g.w+xy.x]})
		}
	}
	dleft := n.d.rot(-1)
	if xy := n.xy.adv(dleft); g.valid(xy) {
		res = append(res, &Node{xy, dleft, 1, n.heatLoss + g.data[xy.y*g.w+xy.x]})
	}

	dright := n.d.rot(+1)
	if xy := n.xy.adv(dright); g.valid(xy) {
		res = append(res, &Node{xy, dright, 1, n.heatLoss + g.data[xy.y*g.w+xy.x]})
	}

	return res
}

func (g Grid) target(n *Node, mindlen int) bool {
	return n.xy.x == g.w-1 && n.xy.y == g.h-1 && n.dlen >= mindlen
}

type Node struct {
	xy       XY
	d        Dir
	dlen     int
	heatLoss int
}

func (n Node) key() string {
	return fmt.Sprintf("%d/%d/%d/%d", n.xy.x, n.xy.y, n.d, n.dlen)
}

func (g Grid) astar(mindlen int, maxdlen int) int {
	nodes := make(map[string]*Node)
	pending := make([]*Node, 0)

	startE := Node{XY{0, 0}, E, 0, 0}
	nodes[startE.key()] = &startE
	pending = append(pending, &startE)

	startS := Node{XY{0, 0}, S, 0, 0}
	nodes[startS.key()] = &startS
	pending = append(pending, &startS)

	minHeatLoss := -1

	for len(pending) > 0 {
		minCost := -1
		minI := 0
		for i, n := range pending {
			cost := n.heatLoss + g.w - n.xy.x + g.h - n.xy.y - 2
			if minCost < 0 || cost < minCost {
				minCost = cost
				minI = i
			}
		}

		n := pending[minI]
		if minI+1 != len(pending) {
			pending[minI] = pending[len(pending)-1]
		}
		pending = pending[:len(pending)-1]

		if g.target(n, mindlen) {
			if minHeatLoss < 0 || n.heatLoss < minHeatLoss {
				minHeatLoss = n.heatLoss
			}
		}

		for _, newN := range g.next(n, mindlen, maxdlen) {
			k := newN.key()
			if existingN, ok := nodes[k]; ok {
				if newN.heatLoss < existingN.heatLoss {
					existingN.heatLoss = newN.heatLoss
					pending = append(pending, existingN)
				}
			} else {
				nodes[k] = newN
				pending = append(pending, newN)
			}
		}
	}

	return minHeatLoss
}

func main() {
	g := &Grid{make([]int, 0), 0, 0}

	helpers.ReadStdin(func(line string) {
		g.add(line)
	})

	if helpers.Part1() {
		fmt.Println(g.astar(0, 3))
	} else {
		fmt.Println(g.astar(4, 10))
	}
}
