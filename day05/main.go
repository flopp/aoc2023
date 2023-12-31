package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

type Range struct {
	start int
	end   int
}

func createRangeFromStartLen(start, length int) Range {
	return Range{start, start + length - 1}
}

func (r *Range) contains(i int) bool {
	return i >= r.start && i <= r.end
}

func (r *Range) overlaps(other Range) bool {
	return r.start <= other.end && other.start <= r.end
}

type RangeMap struct {
	source Range
	offset int
}
type Mapping struct {
	ranges []RangeMap
}

func (m *Mapping) mapValue(source int) int {
	for _, r := range m.ranges {
		if r.source.contains(source) {
			return source + r.offset
		}
	}
	return source
}

func (m *Mapping) mapRange(source Range) []Range {
	// step 1: split source ranges wrt. m.ranges
	split_source := make([]Range, 0)
	split_source = append(split_source, source)
	for i := 0; i < len(split_source); i += 1 {
		s := split_source[i]
		for _, r := range m.ranges {
			// left overlap
			if s.start < r.source.start && s.end >= r.source.start {
				split_source = append(split_source, Range{s.start, r.source.start - 1})
				split_source[i].start = r.source.start
				// do split_source[i] again in the next loop
				i -= 1
				break
			}

			// right overlap
			if s.start <= r.source.end && s.end > r.source.end {
				split_source = append(split_source, Range{r.source.end + 1, s.end})
				split_source[i].end = r.source.end
				// do split_source[i] again in the next loop
				i -= 1
				break
			}
		}
	}

	// step 2: map the split ranges (after splitting, each
	// source range can only overlap with one range from
	// m.ranges)
	mapped := make([]Range, 0)
	for _, s := range split_source {
		mapped_s := s
		for _, r := range m.ranges {
			if s.overlaps(r.source) {
				mapped_s = Range{s.start + r.offset, s.end + r.offset}
				break
			}
		}
		mapped = append(mapped, mapped_s)
	}
	return mapped
}

func main() {
	seeds := make([]int, 0)
	mappings := make([]*Mapping, 0)
	var mapping *Mapping = nil

	helpers.ReadStdin(func(line string) {
		if line == "" {
			// nothing
		} else if strings.HasPrefix(line, "seeds:") {
			for i, seed := range strings.Split(line, " ") {
				if i > 0 {
					seeds = append(seeds, helpers.MustParseInt(seed))
				}
			}
		} else if strings.HasSuffix(line, ":") {
			mapping = &Mapping{make([]RangeMap, 0)}
			mappings = append(mappings, mapping)
		} else {
			a := strings.Split(line, " ")
			dest := helpers.MustParseInt(a[0])
			source := helpers.MustParseInt(a[1])
			length := helpers.MustParseInt(a[2])
			mapping.ranges = append(mapping.ranges, RangeMap{createRangeFromStartLen(source, length), dest - source})
		}
	})

	minLocation := -1
	if helpers.Part1() {
		for _, seed := range seeds {
			value := seed
			for _, mapping := range mappings {
				value = mapping.mapValue(value)
			}
			if minLocation < 0 || value < minLocation {
				minLocation = value
			}
		}
	} else {
		// interpret seeds as ranges
		for i := 0; i < len(seeds); i += 2 {
			values := make([]Range, 0)
			values = append(values, createRangeFromStartLen(seeds[i], seeds[i+1]))

			for _, mapping := range mappings {
				nextValues := make([]Range, 0)
				for _, v := range values {
					nextValues = append(nextValues, mapping.mapRange(v)...)
				}
				values = nextValues
			}

			for _, v := range values {
				if minLocation < 0 || v.start < minLocation {
					minLocation = v.start
				}
			}
		}
	}
	fmt.Println(minLocation)
}
