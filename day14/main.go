package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

type Board struct {
	data [][]byte
	w, h int
}

func createBoard() *Board {
	return &Board{make([][]byte, 0), 0, 0}
}

func (b *Board) add(row string) {
	d := make([]byte, 0, len(row))
	for _, cell := range row {
		d = append(d, byte(cell))
	}
	b.data = append(b.data, d)
	b.h += 1
	b.w = len(row)
}

func (b *Board) copyFrom(other *Board) {
	if b.h != other.h || b.w != other.w {
		b.h = other.h
		b.w = other.w
		b.data = make([][]byte, 0, b.h)
		for y := 0; y < b.h; y += 1 {
			b.data = append(b.data, make([]byte, b.w))
		}
	}
	for y := 0; y < b.h; y += 1 {
		for x := 0; x < b.w; x += 1 {
			b.data[y][x] = other.data[y][x]
		}
	}
}

func (b Board) same(other *Board) bool {
	if b.h != other.h || b.w != other.w {
		return false
	}
	for y := 0; y < b.h; y += 1 {
		for x := 0; x < b.w; x += 1 {
			if b.data[y][x] != other.data[y][x] {
				return false
			}
		}
	}
	return true
}

func (b *Board) tiltV(delta int) {
	from := 0
	to := b.h
	if delta < 0 {
		from = b.h - 1
		to = -1
	}
	for x := 0; x < b.w; x += 1 {
		y0 := -1
		for y := from; y != to; y += delta {
			c := b.data[y][x]
			if y0 == -1 {
				if c == '.' {
					y0 = y
				}
			} else {
				if c == '.' {
				} else if c == '#' {
					y0 = -1
				} else {
					b.data[y0][x] = c
					b.data[y][x] = '.'
					y0 += delta
				}
			}
		}
	}
}

func (b *Board) tiltN() {
	b.tiltV(+1)
}

func (b *Board) tiltS() {
	b.tiltV(-1)
}

func (b *Board) tiltH(delta int) {
	from := 0
	to := b.w
	if delta < 0 {
		from = b.w - 1
		to = -1
	}
	for y := 0; y < b.h; y += 1 {
		x0 := -1
		for x := from; x != to; x += delta {
			c := b.data[y][x]
			if x0 == -1 {
				if c == '.' {
					x0 = x
				}
			} else {
				if c == '.' {
				} else if c == '#' {
					x0 = -1
				} else {
					b.data[y][x0] = c
					b.data[y][x] = '.'
					x0 += delta
				}
			}
		}
	}
}

func (b *Board) tiltW() {
	b.tiltH(+1)
}

func (b *Board) tiltE() {
	b.tiltH(-1)
}

func (b *Board) tilt() {
	b.tiltN()
	b.tiltW()
	b.tiltS()
	b.tiltE()
}

func (b Board) load() int {
	l := 0
	for y, row := range b.data {
		for _, cell := range row {
			if cell == 'O' {
				l += (b.h - y)
			}
		}
	}
	return l
}

func main() {
	b := createBoard()
	helpers.ReadStdin(func(line string) {
		b.add(line)
	})

	if helpers.Part1() {
		b.tiltN()
	} else {
		history := make([]*Board, 0)
		count := 1000000000
		for i := 0; i < count; i += 1 {
			b.tilt()

			found := -1
			for j, other := range history {
				if other.same(b) {
					found = j
					break
				}
			}
			if found >= 0 {
				loop := i - found
				f := (count-found)%loop + found
				b.copyFrom(history[f-1])
				break
			}

			other := createBoard()
			other.copyFrom(b)
			history = append(history, other)
		}
	}
	fmt.Println(b.load())
}
