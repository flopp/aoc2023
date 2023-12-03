package main

import (
	"fmt"

	"github.com/flopp/aoc2023/helpers"
)

func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}
func hasSymbol(lines []string, w, h, minx, maxx, y int) bool {
	for yy := helpers.Max(y-1, 0); yy <= helpers.Min(y+1, h-1); yy += 1 {
		line := lines[yy]
		for xx := helpers.Max(minx-1, 0); xx <= helpers.Min(maxx+1, w-1); xx += 1 {
			if isSymbol(line[xx]) {
				return true
			}
		}
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

	// find start and end of number
	minx := x + dx
	for isDigit(line, minx-1) {
		minx -= 1
	}
	maxx := x + dx
	for isDigit(line, maxx+1) {
		maxx += 1
	}
	return helpers.MustParseInt(line[minx : 1+maxx])
}

func main() {
	lines := make([]string, 0)
	helpers.ReadStdin(func(line string) {
		lines = append(lines, line)
	})
	h := len(lines)
	w := len(lines[0])

	sum := 0
	if helpers.Part1() {
		for y, line := range lines {
			minx := -1
			num := 0
			for x := 0; x < w; x += 1 {
				if isDigit(line, x) {
					if minx == -1 {
						minx = x
						num = 0
					}
					num = 10*num + int(line[x]-'0')
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
