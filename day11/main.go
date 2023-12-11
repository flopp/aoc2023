package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type XY struct {
	x, y int
}

func (xy XY) manhattan(other XY) int {
	dx := xy.x - other.x
	if dx < 0 {
		dx = -dx
	}
	dy := xy.y - other.y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func main() {
	galaxies := make([]XY, 0)

	// parse galaxies, determine empty columns/rows
	emptyX := make([]bool, 0)
	emptyY := make([]bool, 0)
	xy := XY{0, 0}
	helpers.ReadStdin(func(line string) {
		for xy.x = 0; xy.x < len(line); xy.x += 1 {
			if line[xy.x] == '#' {
				galaxies = append(galaxies, xy)

				for xy.x >= len(emptyX) {
					emptyX = append(emptyX, true)
				}
				for xy.y >= len(emptyY) {
					emptyY = append(emptyY, true)
				}
				emptyX[xy.x] = false
				emptyY[xy.y] = false
			}
		}
		xy.y += 1
	})

	// fetch spacing for part
	spacing := 1
	if helpers.Part1() {
		spacing = 2
	} else if helpers.Test() {
		spacing = 10
	} else {
		spacing = 1000000
	}

	// determine coordinates transformation
	x := 0
	newX := make([]int, 0)
	for _, e := range emptyX {
		newX = append(newX, x)
		if e {
			x += spacing
		} else {
			x += 1
		}
	}
	y := 0
	newY := make([]int, 0)
	for _, e := range emptyY {
		newY = append(newY, y)
		if e {
			y += spacing
		} else {
			y += 1
		}
	}

	// apply coordinates transformation
	for i, xy := range galaxies {
		galaxies[i].x = newX[xy.x]
		galaxies[i].y = newY[xy.y]
	}

	// sum distances
	sum := 0
	for i1, xy1 := range galaxies {
		for i2 := i1 + 1; i2 < len(galaxies); i2 += 1 {
			sum += xy1.manhattan(galaxies[i2])
		}
	}

	fmt.Println(sum)
}
