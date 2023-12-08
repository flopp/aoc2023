package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

const (
	Card2 = iota
	Card3 = iota
	Card4 = iota
	Card5 = iota
	Card6 = iota
	Card7 = iota
	Card8 = iota
	Card9 = iota
	CardT = iota
	CardJ = iota
	CardQ = iota
	CardK = iota
	CardA = iota
)
const (
	HandHigh      = iota
	HandOnePair   = iota
	HandTwoPairs  = iota
	HandThree     = iota
	HandFullHouse = iota
	HandFour      = iota
	HandFive      = iota
)

func createCard(c byte) int {
	switch c {
	case '2':
		return Card2
	case '3':
		return Card3
	case '4':
		return Card4
	case '5':
		return Card5
	case '6':
		return Card6
	case '7':
		return Card7
	case '8':
		return Card8
	case '9':
		return Card9
	case 'T':
		return CardT
	case 'J':
		return CardJ
	case 'Q':
		return CardQ
	case 'K':
		return CardK
	case 'A':
		return CardA
	}
	return CardA
}

type Hand [5]int

func createHand(s string) Hand {
	var h Hand
	for i := 0; i < 5; i += 1 {
		h[i] = createCard(s[i])
	}
	return h
}

func handTypeCounts(counts map[int]int) int {
	max := 0
	for _, count := range counts {
		if count > max {
			max = count
		}
	}

	switch len(counts) {
	case 1:
		return HandFive
	case 2:
		if max == 4 {
			return HandFour
		}
		return HandFullHouse
	case 3:
		if max == 3 {
			return HandThree
		}
		return HandTwoPairs
	case 4:
		return HandOnePair
	}

	return HandHigh
}

func handType(hand Hand) int {
	counts := make(map[int]int)
	for i := 0; i < 5; i += 1 {
		c := hand[i]
		if count, ok := counts[c]; ok {
			counts[c] = count + 1
		} else {
			counts[c] = 1
		}
	}

	return handTypeCounts(counts)
}

func handTypeJokers(hand Hand) int {
	jokers := 0
	counts := make(map[int]int)
	for i := 0; i < 5; i += 1 {
		c := hand[i]
		if c == CardJ {
			jokers += 1
		} else if count, ok := counts[c]; ok {
			counts[c] = count + 1
		} else {
			counts[c] = 1
		}
	}

	maxCount := 0
	maxCard := CardJ

	for card, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCard = card
		}
	}

	counts[maxCard] += jokers

	return handTypeCounts(counts)
}

type HandBid struct {
	h Hand
	t int
	b int
}

func (hb1 *HandBid) isLessThan(hb2 *HandBid) bool {
	if hb1.t < hb2.t {
		return true
	}
	if hb1.t > hb2.t {
		return false
	}

	for i := 0; i < 5; i += 1 {
		c1 := hb1.h[i]
		c2 := hb2.h[i]
		if c1 < c2 {
			return true
		}
		if c1 > c2 {
			return false
		}
	}
	return false
}

func (hb1 *HandBid) isLessThanJokers(hb2 *HandBid) bool {
	if hb1.t < hb2.t {
		return true
	}
	if hb1.t > hb2.t {
		return false
	}

	for i := 0; i < 5; i += 1 {
		c1 := hb1.h[i]
		c2 := hb2.h[i]
		if c1 == CardJ && c2 != CardJ {
			return true
		}
		if c1 != CardJ && c2 == CardJ {
			return false
		}
		if c1 < c2 {
			return true
		}
		if c1 > c2 {
			return false
		}
	}
	return false
}

func main() {
	handBids := make([]*HandBid, 0)
	helpers.ReadStdin(func(line string) {
		a := strings.Split(line, " ")
		h := createHand(a[0])
		var t int
		if helpers.Part1() {
			t = handType(h)
		} else {
			t = handTypeJokers(h)
		}
		handBids = append(handBids, &HandBid{h, t, helpers.MustParseInt(a[1])})
	})

	if helpers.Part1() {
		sort.Slice(handBids, func(i, j int) bool {
			return handBids[i].isLessThan(handBids[j])
		})
	} else {
		sort.Slice(handBids, func(i, j int) bool {
			return handBids[i].isLessThanJokers(handBids[j])
		})
	}

	winnings := 0
	for i, handBid := range handBids {
		winnings += (i + 1) * handBid.b
	}
	fmt.Println(winnings)
}
