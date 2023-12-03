package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

func getSymbol(lines []string, w, h, x, y int) byte {
	if x >= 0 && x < w && y >= 0 && y < h {
		return lines[y][x]
	}
	return '.'
}
func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}
func hasSymbol(lines []string, w, h, minx, maxx, y int) bool {
	// top + bottom
	for x := minx - 1; x <= maxx+1; x += 1 {
		if isSymbol(getSymbol(lines, w, h, x, y-1)) {
			return true
		}
		if isSymbol(getSymbol(lines, w, h, x, y+1)) {
			return true
		}
	}
	// left
	if isSymbol(getSymbol(lines, w, h, minx-1, y)) {
		return true
	}
	// right
	if isSymbol(getSymbol(lines, w, h, maxx+1, y)) {
		return true
	}

	return false
}

func isDigit(line string, x int) bool {
	if x < 0 || x >= len(line) {
		return false
	}
	c := line[x]
	return '0' <= c && c <= '9'
}

func getNum(lines []string, w, h, x, y, dx, dy int) int {
	if x+dx < 0 || x+dx >= w || y+dy < 0 || y+dy >= h {
		return -1
	}
	line := lines[y+dy]
	if !isDigit(line, x+dx) {
		return -1
	}

	minx := x + dx
	for isDigit(line, minx-1) {
		minx -= 1
	}
	maxx := x + dx
	for isDigit(line, maxx+1) {
		maxx += 1
	}
	s := line[minx : 1+maxx]
	return helpers.MustParseInt(s)
}

func main() {
	w := -1
	lines := make([]string, 0)
	helpers.ReadStdin(func(line string) {
		if w == -1 {
			w = len(line)
		} else if w != len(line) {
			panic(fmt.Errorf("bad line: len=%d w=%d", len(line), w))
		}
		lines = append(lines, line)
	})
	h := len(lines)
	if h == 0 {
		panic("no lines")
	}

	sum := 0
	if helpers.Part1() {
		for y, line := range lines {
			minx := -1
			num := 0
			for x := 0; x < w; x += 1 {
				c := line[x]
				if '0' <= c && c <= '9' {
					if minx == -1 {
						minx = x
						num = 0
					}
					num = 10*num + int(c-'0')
				} else if minx != -1 {
					if hasSymbol(lines, w, h, minx, x-1, y) {
						sum += num
					}
					minx = -1
				}
			}
			if minx != -1 {
				if hasSymbol(lines, w, h, minx, w-1, y) {
					sum += num
				}
				minx = -1
			}
		}
	} else {
		for y, line := range lines {
			for x := 0; x < w; x += 1 {
				if line[x] != '*' {
					// skip non-gears
					continue
				}
				mult := 1
				nums := 0
				for dy := -1; dy <= +1; dy += 1 {
					n := getNum(lines, w, h, x, y, 0, dy)
					if n != -1 {
						nums += 1
						mult *= n
					} else {
						for dx := -1; dx <= +1; dx += 2 {
							n = getNum(lines, w, h, x, y, dx, dy)
							if n != -1 {
								nums += 1
								mult *= n
							}
						}
					}
				}
				if nums == 2 {
					sum += mult
				}
			}
		}
	}
	fmt.Println(sum)
}
