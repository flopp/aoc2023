package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

type Draw struct {
	r, g, b int
}

func (draw *Draw) isPossible(maxR, maxG, maxB int) bool {
	return draw.r <= maxR && draw.g <= maxG && draw.b <= maxB
}

type Game struct {
	id    int
	draws []*Draw
}

func (game *Game) isPossible(maxR, maxG, maxB int) bool {
	for _, d := range game.draws {
		if !d.isPossible(maxR, maxG, maxB) {
			return false
		}
	}
	return true
}

func (game *Game) power() int {
	r := 0
	g := 0
	b := 0
	for _, d := range game.draws {
		r = helpers.Max(r, d.r)
		g = helpers.Max(g, d.g)
		b = helpers.Max(b, d.b)
	}
	return r * g * b
}

var re_game = regexp.MustCompile(`^Game (\d+): (.*)$`)
var re_cubes = regexp.MustCompile(`^(\d+) (red|green|blue)$`)

func parseGame(line string) *Game {
	mg := re_game.FindStringSubmatch(line)
	if mg == nil {
		return nil
	}
	g := Game{helpers.MustParseInt(mg[1]), make([]*Draw, 0)}
	for _, s := range strings.Split(mg[2], "; ") {
		d := Draw{0, 0, 0}
		for _, cubes := range strings.Split(s, ", ") {
			mc := re_cubes.FindStringSubmatch(cubes)
			if mc == nil {
				return nil
			}
			switch mc[2] {
			case "red":
				d.r = helpers.MustParseInt(mc[1])
			case "green":
				d.g = helpers.MustParseInt(mc[1])
			case "blue":
				d.b = helpers.MustParseInt(mc[1])
			}
			g.draws = append(g.draws, &d)
		}
	}
	return &g
}

func main() {
	sum := 0
	helpers.ReadStdin(func(line string) {
		game := parseGame(line)
		if game == nil {
			panic(fmt.Errorf("bad line: %s", line))
		}
		if helpers.Part1() {
			if game.isPossible(12, 13, 14) {
				sum += game.id
			}
		} else {
			sum += game.power()
		}
	})
	fmt.Println(sum)
}
