package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

const (
	N = 1
	E = 2
	S = 4
	W = 8
)

type XY struct {
	x, y int
}

type Ray struct {
	xy  XY
	dir int
}

type MirrorGrid struct {
	grid      []byte
	energized []int
	w, h      int
	rays      []Ray
}

func createMirrorGrid() *MirrorGrid {
	return &MirrorGrid{make([]byte, 0), make([]int, 0), 0, 0, make([]Ray, 0)}
}

func (mg *MirrorGrid) add(line string) {
	mg.h += 1
	if mg.w == 0 {
		mg.w = len(line)
	}
	for _, c := range line {
		mg.grid = append(mg.grid, byte(c))
		mg.energized = append(mg.energized, 0)
	}
}

func (mg *MirrorGrid) countEnergized() int {
	c := 0
	for _, e := range mg.energized {
		if e != 0 {
			c += 1
		}
	}
	return c
}

func (mg *MirrorGrid) setEnergized(xy XY, dir int) bool {
	e := mg.energized[xy.y*mg.w+xy.x]
	if (e & dir) == dir {
		return false
	}
	mg.energized[xy.y*mg.w+xy.x] = e | dir
	return true
}

func (mg *MirrorGrid) resetEnergized() {
	for i := range mg.energized {
		mg.energized[i] = 0
	}
}

func (mg *MirrorGrid) addRay(xy XY, dir int) {
	if mg.setEnergized(xy, dir) {
		mg.rays = append(mg.rays, Ray{xy, dir})
	}
}

func (mg *MirrorGrid) simulate() {
	nr0 := Ray{XY{}, 0}
	nr1 := Ray{XY{}, 0}
	for len(mg.rays) > 0 {
		r := mg.rays[len(mg.rays)-1]
		mg.rays = mg.rays[:len(mg.rays)-1]

		nr0.dir = 0
		nr0.xy = r.xy
		nr1.dir = 0
		nr1.xy = r.xy

		switch mg.grid[r.xy.y*mg.w+r.xy.x] {
		case '.':
			switch r.dir {
			case N:
				nr0.dir = N
				nr0.xy.y -= 1
			case S:
				nr0.dir = S
				nr0.xy.y += 1
			case E:
				nr0.dir = E
				nr0.xy.x += 1
			case W:
				nr0.dir = W
				nr0.xy.x -= 1
			}
		case '/':
			switch r.dir {
			case N:
				nr0.dir = E
				nr0.xy.x += 1
			case S:
				nr0.dir = W
				nr0.xy.x -= 1
			case E:
				nr0.dir = N
				nr0.xy.y -= 1
			case W:
				nr0.dir = S
				nr0.xy.y += 1
			}
		case '\\':
			switch r.dir {
			case N:
				nr0.dir = W
				nr0.xy.x -= 1
			case S:
				nr0.dir = E
				nr0.xy.x += 1
			case E:
				nr0.dir = S
				nr0.xy.y += 1
			case W:
				nr0.dir = N
				nr0.xy.y -= 1
			}
		case '-':
			switch r.dir {
			case N:
				nr0.dir = W
				nr0.xy.x -= 1
				nr1.dir = E
				nr1.xy.x += 1
			case S:
				nr0.dir = W
				nr0.xy.x -= 1
				nr1.dir = E
				nr1.xy.x += 1
			case E:
				nr0.dir = E
				nr0.xy.x += 1
			case W:
				nr0.dir = W
				nr0.xy.x -= 1
			}
		case '|':
			switch r.dir {
			case E:
				nr0.dir = N
				nr0.xy.y -= 1
				nr1.dir = S
				nr1.xy.y += 1
			case W:
				nr0.dir = N
				nr0.xy.y -= 1
				nr1.dir = S
				nr1.xy.y += 1
			case N:
				nr0.dir = N
				nr0.xy.y -= 1
			case S:
				nr0.dir = S
				nr0.xy.y += 1
			}
		}

		if nr0.dir != 0 && nr0.xy.x >= 0 && nr0.xy.x < mg.w && nr0.xy.y >= 0 && nr0.xy.y < mg.h {
			mg.addRay(nr0.xy, nr0.dir)
		}
		if nr1.dir != 0 && nr1.xy.x >= 0 && nr1.xy.x < mg.w && nr1.xy.y >= 0 && nr1.xy.y < mg.h {
			mg.addRay(nr1.xy, nr1.dir)
		}
	}
}

func main() {
	mg := createMirrorGrid()
	helpers.ReadStdin(func(line string) {
		mg.add(line)
	})

	if helpers.Part1() {
		mg.addRay(XY{0, 0}, E)
		mg.simulate()
		fmt.Println(mg.countEnergized())
	} else {
		maxEnergized := 0
		for x := 0; x < mg.w; x += 1 {
			mg.resetEnergized()
			mg.addRay(XY{x, 0}, S)
			mg.simulate()
			maxEnergized = helpers.Max(maxEnergized, mg.countEnergized())

			mg.resetEnergized()
			mg.addRay(XY{x, mg.h - 1}, N)
			mg.simulate()
			maxEnergized = helpers.Max(maxEnergized, mg.countEnergized())
		}
		for y := 0; y < mg.h; y += 1 {
			mg.resetEnergized()
			mg.addRay(XY{0, y}, E)
			mg.simulate()
			maxEnergized = helpers.Max(maxEnergized, mg.countEnergized())

			mg.resetEnergized()
			mg.addRay(XY{mg.w - 1, y}, W)
			mg.simulate()
			maxEnergized = helpers.Max(maxEnergized, mg.countEnergized())
		}
		fmt.Println(maxEnergized)
	}
}
