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

func rotate(d, offset int) int {
	return (4 + d + offset) % 4
}

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

type XYD struct {
	xy  XY
	dir int
}

func main() {
	// read-in maze, add '.' border to simplify neighbr checking, etc.
	maze := createMaze()
	helpers.ReadStdin(func(line string) {
		maze.addRow([]byte(line))
	})
	maze.addEmptyRow()

	// determine start position, replace it with matching pipe
	start := maze.start()

	// determine start direction
	dir := -1
	for d := DirN; d <= DirW; d += 1 {
		if isOpen(maze.get(start), d) {
			dir = d
			break
		}
	}

	// follow pipes until we hit start again
	steps := make([]XYD, 0)
	steps = append(steps, XYD{start, dir})
	for {
		xy := steps[len(steps)-1].xy.next(dir)
		if xy == start {
			break
		}
		pipe := maze.get(xy)
		dir = nextD(dir, pipe)
		steps = append(steps, XYD{xy, dir})
	}

	if helpers.Part1() {
		fmt.Println(len(steps) / 2)
	} else {
		// remove non-path pipes
		xys := make([]XY, 0, len(steps))
		for _, xyd := range steps {
			xys = append(xys, xyd.xy)
		}
		maze.clean(xys)

		// fill border region
		maze.floodFill(XY{0, 0}, 'O')

		// determine if 'O' is left or right of path
		outerRot := 0
		for _, xyd := range steps {
			if maze.get(xyd.xy.next(rotate(xyd.dir, -1))) == 'O' {
				outerRot = -1
				break
			} else if maze.get(xyd.xy.next(rotate(xyd.dir, +1))) == 'O' {
				outerRot = +1
				break
			}
		}
		if outerRot == 0 {
			panic("cannot find outer side")
		}

		// follow path and fill outer side
		for i, xyd := range steps {
			d := rotate(xyd.dir, outerRot)
			maze.floodFill(xyd.xy.next(d), 'O')
			maze.floodFill(steps[(i+1)%len(steps)].xy.next(d), 'O')
		}

		fmt.Println(maze.count('.'))
	}
}
