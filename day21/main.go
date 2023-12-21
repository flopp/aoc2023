package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

func createEmptyRow(w int) []byte {
	row := make([]byte, w)
	for i := range row {
		row[i] = '.'
	}
	return row
}
func createRow(offset int, line string) []byte {
	row := make([]byte, 0, 2*offset+len(line))
	for i := 0; i < offset; i += 1 {
		row = append(row, '.')
	}
	for _, c := range line {
		row = append(row, byte(c))
	}
	for i := 0; i < offset; i += 1 {
		row = append(row, '.')
	}
	return row
}

type XY struct {
	x, y int
}

type Item struct {
	xy    XY
	steps int
}

func (i Item) key() string {
	return fmt.Sprintf("%d-%d-%d", i.xy.x, i.xy.y, i.steps)
}

func main() {
	steps := 6
	if helpers.Puzzle() {
		steps = 64
	}

	w := 0
	grid := make([][]byte, 0)
	helpers.ReadStdin(func(line string) {
		if w == 0 {
			w = 2*steps + len(line)
			for y := 0; y < steps; y += 1 {
				grid = append(grid, createEmptyRow(w))
			}
		}
		grid = append(grid, createRow(steps, line))
	})
	for y := 0; y < steps; y += 1 {
		grid = append(grid, createEmptyRow(w))
	}
	h := len(grid)

	start := &Item{XY{-1, -1}, steps}
	for y := 0; y < h; y += 1 {
		row := grid[y]
		for x := 0; x < w; x += 1 {
			if row[x] == 'S' {
				start.xy.x = x
				start.xy.y = y
			}
		}
	}

	seen := make(map[string]*Item)
	pending := make([]*Item, 0)

	seen[start.key()] = start
	pending = append(pending, start)
	for len(pending) > 0 {
		item := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if item.steps == 0 {
			continue
		}

		// north
		if grid[item.xy.y-1][item.xy.x] != '#' {
			next := &Item{XY{item.xy.x, item.xy.y - 1}, item.steps - 1}
			if _, ok := seen[next.key()]; !ok {
				seen[next.key()] = next
				pending = append(pending, next)
			}
		}
		// south
		if grid[item.xy.y+1][item.xy.x] != '#' {
			next := &Item{XY{item.xy.x, item.xy.y + 1}, item.steps - 1}
			if _, ok := seen[next.key()]; !ok {
				seen[next.key()] = next
				pending = append(pending, next)
			}
		}
		// east
		if grid[item.xy.y][item.xy.x+1] != '#' {
			next := &Item{XY{item.xy.x + 1, item.xy.y}, item.steps - 1}
			if _, ok := seen[next.key()]; !ok {
				seen[next.key()] = next
				pending = append(pending, next)
			}
		}
		// west
		if grid[item.xy.y][item.xy.x-1] != '#' {
			next := &Item{XY{item.xy.x - 1, item.xy.y}, item.steps - 1}
			if _, ok := seen[next.key()]; !ok {
				seen[next.key()] = next
				pending = append(pending, next)
			}
		}
	}

	count := 0
	for _, item := range seen {
		if item.steps == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
