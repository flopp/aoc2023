package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2023/helpers"
)

type XY struct {
	x, y int
}

func (xy XY) adv(d Dir, dist int) XY {
	switch d {
	case U:
		return XY{xy.x, xy.y - dist}
	case R:
		return XY{xy.x + dist, xy.y}
	case D:
		return XY{xy.x, xy.y + dist}
	case L:
		return XY{xy.x - dist, xy.y}
	}
	return xy
}

type Range struct {
	min, max int
}

func (r *Range) update(value int) {
	if value < r.min {
		r.min = value
	} else if value > r.max {
		r.max = value
	}
}

type Dir int

const (
	U = Dir(0)
	R = Dir(1)
	D = Dir(2)
	L = Dir(3)
)

func String2Dir(s string) Dir {
	switch s {
	case "U":
		return U
	case "R":
		return R
	case "D":
		return D
	case "L":
		return L
	}
	panic(fmt.Errorf("bad dir: %s", s))
}

type Dig struct {
	dir   Dir
	dist  int
	color string
}

func dig(grid []byte, rx, ry Range, digs []*Dig) {
	w := 1 + rx.max - rx.min
	xy := XY{0, 0}
	grid[(xy.y-ry.min)*w+xy.x-rx.min] = '#'
	for _, dig := range digs {
		for i := 0; i < dig.dist; i += 1 {
			xy = xy.adv(dig.dir, 1)
			grid[(xy.y-ry.min)*w+xy.x-rx.min] = '#'
		}
	}
}

func fill(grid []byte, rx, ry Range) {
	h := 1 + ry.max - ry.min
	w := 1 + rx.max - rx.min
	pending := make([]XY, 0)
	for x := 0; x < w; x += 1 {
		if grid[0*w+x] == '.' {
			pending = append(pending, XY{x, 0})
		}
		if grid[(h-1)*w+x] == '.' {
			pending = append(pending, XY{x, h - 1})
		}
	}
	for y := 0; y < h; y += 1 {
		if grid[y*w+0] == '.' {
			pending = append(pending, XY{0, y})
		}
		if grid[y*w+w-1] == '.' {
			pending = append(pending, XY{w - 1, y})
		}
	}

	for len(pending) > 0 {
		xy := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if grid[xy.y*w+xy.x] == '.' {
			grid[xy.y*w+xy.x] = '!'
			if xy.x > 0 {
				pending = append(pending, xy.adv(L, 1))
			}
			if xy.x < w-1 {
				pending = append(pending, xy.adv(R, 1))
			}
			if xy.y > 0 {
				pending = append(pending, xy.adv(U, 1))
			}
			if xy.y < h-1 {
				pending = append(pending, xy.adv(D, 1))
			}
		}
	}
	for i := range grid {
		if grid[i] == '!' {
			grid[i] = '.'
		} else if grid[i] == '.' {
			grid[i] = '#'
		}
	}
}

func count(grid []byte, rx, ry Range) int {
	n := 0
	for _, c := range grid {
		if c == '#' {
			n += 1
		}
	}
	return n
}

func main() {
	if helpers.Part1() {
		digs := make([]*Dig, 0)
		re_dig := regexp.MustCompile(`^(.) (\d+) \(#(.*)\)$`)
		helpers.ReadStdin(func(line string) {
			m := re_dig.FindStringSubmatch(line)
			if m != nil {
				digs = append(digs, &Dig{String2Dir(m[1]), helpers.MustParseInt(m[2]), m[3]})
			}
		})

		// determine ranges
		rx := Range{0, 0}
		ry := Range{0, 0}
		xy := XY{0, 0}
		for _, dig := range digs {
			xy = xy.adv(dig.dir, dig.dist)
			rx.update(xy.x)
			ry.update(xy.y)
		}

		w := 1 + rx.max - rx.min
		h := 1 + ry.max - ry.min
		grid := make([]byte, w*h)
		for i := range grid {
			grid[i] = '.'
		}

		dig(grid, rx, ry, digs)
		fill(grid, rx, ry)
		fmt.Println(count(grid, rx, ry))
	}
}
