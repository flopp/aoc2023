package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

const (
	DirN = iota
	DirE = iota
	DirS = iota
	DirW = iota
)

type XY struct {
	x, y int
}

func (xy XY) N() XY  { return XY{xy.x, xy.y - 1} }
func (xy XY) NE() XY { return XY{xy.x + 1, xy.y - 1} }
func (xy XY) NW() XY { return XY{xy.x - 1, xy.y - 1} }
func (xy XY) E() XY  { return XY{xy.x + 1, xy.y} }
func (xy XY) S() XY  { return XY{xy.x, xy.y + 1} }
func (xy XY) SE() XY { return XY{xy.x + 1, xy.y + 1} }
func (xy XY) SW() XY { return XY{xy.x - 1, xy.y + 1} }
func (xy XY) W() XY  { return XY{xy.x - 1, xy.y} }

func (xy XY) next(d int) XY {
	switch d {
	case DirN:
		return xy.N()
	case DirE:
		return xy.E()
	case DirS:
		return xy.S()
	case DirW:
		return xy.W()
	}
	panic(fmt.Errorf("bad d: %d", d))
}

func (xy XY) dir(next XY) int {
	if next.x == xy.x {
		if next.y == xy.y-1 {
			return DirN
		}
		if next.y == xy.y+1 {
			return DirS
		}
	}
	if next.y == xy.y {
		if next.x == xy.x-1 {
			return DirW
		}
		if next.x == xy.x+1 {
			return DirE
		}
	}
	panic(fmt.Errorf("dir: %v => %v", xy, next))
}

func nextD(d int, pipe byte) int {
	switch pipe {
	case '|':
		switch d {
		case DirN:
			return d
		case DirS:
			return d
		}
	case '-':
		switch d {
		case DirE:
			return d
		case DirW:
			return d
		}
	case 'L':
		switch d {
		case DirS:
			return DirE
		case DirW:
			return DirN
		}
	case 'F':
		switch d {
		case DirN:
			return DirE
		case DirW:
			return DirS
		}
	case '7':
		switch d {
		case DirN:
			return DirW
		case DirE:
			return DirS
		}
	case 'J':
		switch d {
		case DirS:
			return DirW
		case DirE:
			return DirN
		}
	}
	panic(fmt.Errorf("nextD: bad dir: %d, bad pipe: %v", d, string(rune(pipe))))
}

func isOpen(c byte, d int) bool {
	switch d {
	case DirN:
		return c == 'J' || c == '|' || c == 'L'
	case DirE:
		return c == 'L' || c == '-' || c == 'F'
	case DirS:
		return c == '7' || c == '|' || c == 'F'
	case DirW:
		return c == 'J' || c == '-' || c == '7'
	}
	panic(fmt.Errorf("isOpen: bad dir %d", d))
}

type Maze struct {
	grid [][]byte
	w, h int
	s    XY
}

func createMaze() *Maze {
	return &Maze{make([][]byte, 0), 0, 0, XY{-1, -1}}
}

func (m *Maze) addRow(row []byte) {
	if len(m.grid) == 0 {
		m.w = len(row) + 2
		m.addEmptyRow()
	}
	r := make([]byte, m.w)
	r[0] = '.'
	r[m.w-1] = '.'
	for i, b := range row {
		r[i+1] = b
	}
	m.grid = append(m.grid, r)
	m.h += 1
}

func (m *Maze) addEmptyRow() {
	r := make([]byte, m.w)
	for i := 0; i < m.w; i += 1 {
		r[i] = '.'
	}
	m.grid = append(m.grid, r)
	m.h += 1
}

func (m *Maze) clean(steps []XY) {
	grid := m.grid
	m.grid = make([][]byte, 0, len(grid))
	m.h = 0
	for i := 0; i < len(grid); i += 1 {
		m.addEmptyRow()
	}
	for _, xy := range steps {
		m.grid[xy.y][xy.x] = grid[xy.y][xy.x]
	}
}

func (m Maze) count(c byte) int {
	n := 0
	for _, row := range m.grid {
		for _, c := range row {
			if c == '.' {
				n += 1
			}
		}
	}
	return n
}

func (m *Maze) floodFill(xy XY, fillc byte) {
	pending := make([]XY, 0)
	pending = append(pending, xy)

	for len(pending) > 0 {
		xy = pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		c := m.get(xy)
		if c != '.' {
			continue
		}
		m.set(xy, fillc)
		if xy.x > 0 {
			pending = append(pending, xy.W())
			if xy.y > 0 {
				pending = append(pending, xy.NW())
			}
			if xy.y+1 < m.h {
				pending = append(pending, xy.SW())
			}
		}
		if xy.x+1 < m.w {
			pending = append(pending, xy.E())
			if xy.y > 0 {
				pending = append(pending, xy.NE())
			}
			if xy.y+1 < m.h {
				pending = append(pending, xy.SE())
			}
		}
		if xy.y > 0 {
			pending = append(pending, xy.N())
		}
		if xy.y+1 < m.h {
			pending = append(pending, xy.S())
		}
	}
}

func (m Maze) get(xy XY) byte {
	return m.grid[xy.y][xy.x]
}

func (m *Maze) set(xy XY, c byte) {
	m.grid[xy.y][xy.x] = c
}

func (m *Maze) start() XY {
	if m.s.x >= 0 {
		return m.s
	}
	var xy XY
	for xy.y = 0; xy.y < m.h; xy.y += 1 {
		for xy.x = 0; xy.x < m.w; xy.x += 1 {
			if m.get(xy) == 'S' {
				n := isOpen(m.get(xy.N()), DirS)
				e := isOpen(m.get(xy.E()), DirW)
				s := isOpen(m.get(xy.S()), DirN)
				w := isOpen(m.get(xy.W()), DirE)
				if n {
					if e {
						m.set(xy, 'L')
					} else if s {
						m.set(xy, '|')
					} else if w {
						m.set(xy, 'J')
					} else {
						panic("cannot determine start N")
					}
				} else if e {
					if s {
						m.set(xy, 'F')
					} else if w {
						m.set(xy, '-')
					} else {
						panic("cannot determine start E")
					}
				} else if s {
					if w {
						m.set(xy, '7')
					} else {
						panic("cannot determine start S")
					}
				} else {
					panic("cannot determine start W")
				}
				m.s = xy
				return m.s
			}
		}
	}
	panic("cannot find start in maze")
}

func main() {
	maze := createMaze()
	helpers.ReadStdin(func(line string) {
		maze.addRow([]byte(line))
	})
	maze.addEmptyRow()

	d := -1
	start := maze.start()

	if isOpen(maze.get(start.next(DirN)), DirS) {
		d = DirN
	} else if isOpen(maze.get(start.next(DirE)), DirW) {
		d = DirE
	} else if isOpen(maze.get(start.next(DirS)), DirN) {
		d = DirS
	} else if isOpen(maze.get(start.next(DirW)), DirE) {
		d = DirW
	}

	steps := make([]XY, 0)
	steps = append(steps, start)
	for {
		xy := steps[len(steps)-1].next(d)
		if xy == start {
			break
		}
		pipe := maze.get(xy)
		d = nextD(d, pipe)
		steps = append(steps, xy)
	}

	if helpers.Part1() {
		fmt.Println(len(steps) / 2)
	} else {
		maze.clean(steps)
		maze.floodFill(XY{0, 0}, 'O')

		outer := false
		left := false
		var last XY
		for i := 0; i <= 2*len(steps); i += 1 {
			xy := steps[i%len(steps)]
			if i > 0 {
				switch last.dir(xy) {
				case DirN:
					if !outer {
						if maze.get(xy.W()) == 'O' {
							outer = true
							left = true
						} else if maze.get(xy.E()) == 'O' {
							outer = true
							left = false
						}
					}
				case DirE:
					if !outer {
						if maze.get(xy.N()) == 'O' {
							outer = true
							left = true
						} else if maze.get(xy.S()) == 'O' {
							outer = true
							left = false
						}
					}
				case DirS:
					if !outer {
						if maze.get(xy.E()) == 'O' {
							outer = true
							left = true
						} else if maze.get(xy.W()) == 'O' {
							outer = true
							left = false
						}
					}
				case DirW:
					if !outer {
						if maze.get(xy.S()) == 'O' {
							outer = true
							left = true
						} else if maze.get(xy.N()) == 'O' {
							outer = true
							left = false
						}
					}
				}
			}
			last = xy
		}
		if !outer {
			panic("cannot find outer side")
		}

		for i := 0; i <= 2*len(steps); i += 1 {
			xy := steps[i%len(steps)]
			if i > 0 {
				switch last.dir(xy) {
				case DirN:
					if left {
						maze.floodFill(last.W(), 'O')
						maze.floodFill(xy.W(), 'O')
					} else {
						maze.floodFill(last.E(), 'O')
						maze.floodFill(xy.E(), 'O')
					}
				case DirE:
					if left {
						maze.floodFill(last.N(), 'O')
						maze.floodFill(xy.N(), 'O')
					} else {
						maze.floodFill(last.S(), 'O')
						maze.floodFill(xy.S(), 'O')
					}
				case DirS:
					if left {
						maze.floodFill(last.E(), 'O')
						maze.floodFill(xy.E(), 'O')
					} else {
						maze.floodFill(last.W(), 'O')
						maze.floodFill(xy.W(), 'O')
					}
				case DirW:
					if left {
						maze.floodFill(last.S(), 'O')
						maze.floodFill(xy.S(), 'O')
					} else {
						maze.floodFill(last.N(), 'O')
						maze.floodFill(xy.N(), 'O')
					}
				}
			}
			last = xy
		}

		fmt.Println(maze.count('.'))
	}
}
