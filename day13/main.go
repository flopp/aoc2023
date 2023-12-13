package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type Pattern struct {
	grid [][]byte
}

func (p Pattern) rowDiff(y0, y1 int) int {
	smudges := 0
	w := len(p.grid[0])
	for x := 0; x < w; x += 1 {
		if p.grid[y0][x] != p.grid[y1][x] {
			smudges += 1
		}
	}
	return smudges
}

func (p Pattern) colDiff(x0, x1 int) int {
	smudges := 0
	h := len(p.grid)
	for y := 0; y < h; y += 1 {
		if p.grid[y][x0] != p.grid[y][x1] {
			smudges += 1
		}
	}
	return smudges
}

func (p Pattern) reflects(neededSmudges int, dim int, rowColDiff func(int, int) int) int {
	// we're looking at every row/column twice => smudges are detected twice
	neededSmudges *= 2
	for candidate := 1; candidate < dim; candidate += 1 {
		smudges := 0
		for i := 0; i < dim; i += 1 {
			otheri := 2*candidate - i - 1
			if otheri < 0 || otheri >= dim {
				continue
			}
			smudges += rowColDiff(i, otheri)
			if smudges > neededSmudges {
				break
			}
		}
		if smudges == neededSmudges {
			return candidate
		}
	}
	return 0
}

func (p Pattern) reflectsVertically(neededSmudges int) int {
	return p.reflects(neededSmudges, len(p.grid), p.rowDiff)
}

func (p Pattern) reflectsHorizontally(neededSmudges int) int {
	return p.reflects(neededSmudges, len(p.grid[0]), p.colDiff)
}

func main() {
	patterns := make([]*Pattern, 0)
	pattern := &Pattern{make([][]byte, 0)}
	patterns = append(patterns, pattern)
	helpers.ReadStdin(func(line string) {
		if line == "" {
			pattern = &Pattern{make([][]byte, 0)}
			patterns = append(patterns, pattern)
		} else {
			row := make([]byte, 0, len(line))
			for i := 0; i < len(line); i += 1 {
				row = append(row, line[i])
			}
			pattern.grid = append(pattern.grid, row)
		}
	})

	sum := 0
	neededSmudges := 0
	if !helpers.Part1() {
		neededSmudges = 1
	}
	for _, pattern := range patterns {
		if column := pattern.reflectsHorizontally(neededSmudges); column != 0 {
			sum += column
		} else if row := pattern.reflectsVertically(neededSmudges); row != 0 {
			sum += 100 * row
		} else {
			panic("no reflection")
		}
	}
	fmt.Println(sum)
}
