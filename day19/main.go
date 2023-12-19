package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2023/helpers"
)

type Rule struct {
	category byte
	lt       bool
	value    int
	target   string
}

func (r Rule) accept(part Part) bool {
	if r.lt {
		switch r.category {
		case 'x':
			return part.x < r.value
		case 'm':
			return part.m < r.value
		case 'a':
			return part.a < r.value
		case 's':
			return part.s < r.value
		}
	} else {
		switch r.category {
		case 'x':
			return part.x > r.value
		case 'm':
			return part.m > r.value
		case 'a':
			return part.a > r.value
		case 's':
			return part.s > r.value
		}
	}
	return false
}

type Workflow struct {
	rules  []Rule
	target string
}

func (wf Workflow) next(part Part) string {
	//fmt.Printf("wf=%v part=%v\n", wf, part)
	for _, r := range wf.rules {
		if r.accept(part) {
			//fmt.Printf("    r=%v => %s\n", r, r.target)
			return r.target
		}
	}
	//fmt.Printf("    => %s\n", wf.target)
	return wf.target
}

type Part struct {
	x, m, a, s int
}

func accepted(part Part, workflows map[string]Workflow) bool {
	state := "in"
	for {
		wf, ok := workflows[state]
		if !ok {
			panic(fmt.Errorf("bad name: %s", state))
		}

		state = wf.next(part)
		if state == "A" {
			return true
		} else if state == "R" {
			return false
		}
	}
}

type Range struct {
	min, max int
}

func (r Range) valid() bool {
	return r.min >= 0
}

func (r Range) count() int64 {
	if r.valid() {
		return int64(1 + r.max - r.min)
	}
	return 0
}

type PartRange struct {
	x, m, a, s Range
}

func (r PartRange) count() int64 {
	return r.x.count() * r.m.count() * r.a.count() * r.s.count()
}

func (r PartRange) valid() bool {
	return r.x.valid() && r.m.valid() && r.a.valid() && r.s.valid()
}

func splitRange(r Range, rule Rule) (Range, Range) {
	if rule.lt {
		if rule.value <= r.min {
			return Range{-1, -1}, r
		}
		if rule.value <= r.max {
			return Range{r.min, rule.value - 1}, Range{rule.value, r.max}
		}
		return r, Range{-1, -1}
	}

	if rule.value >= r.max {
		return Range{-1, -1}, r
	}
	if rule.value >= r.min {
		return Range{rule.value + 1, r.max}, Range{r.min, rule.value}
	}
	return r, Range{-1, -1}

}

func split(r PartRange, rule Rule) (PartRange, PartRange) {
	switch rule.category {
	case 'x':
		r1, r2 := splitRange(r.x, rule)
		return PartRange{r1, r.m, r.a, r.s}, PartRange{r2, r.m, r.a, r.s}
	case 'm':
		r1, r2 := splitRange(r.m, rule)
		return PartRange{r.x, r1, r.a, r.s}, PartRange{r.x, r2, r.a, r.s}
	case 'a':
		r1, r2 := splitRange(r.a, rule)
		return PartRange{r.x, r.m, r1, r.s}, PartRange{r.x, r.m, r2, r.s}
	case 's':
		r1, r2 := splitRange(r.s, rule)
		return PartRange{r.x, r.m, r.a, r1}, PartRange{r.x, r.m, r.a, r2}
	}
	panic("bad rule category")
}

func createPartRange(min, max int) PartRange {
	return PartRange{Range{min, max}, Range{min, max}, Range{min, max}, Range{min, max}}
}

func propagate(state string, start PartRange, workflows map[string]Workflow) []PartRange {
	if state == "A" {
		rs := make([]PartRange, 1)
		rs[0] = start
		return rs
	}
	if state == "R" {
		return make([]PartRange, 0)
	}

	wf, ok := workflows[state]
	if !ok {
		panic(fmt.Errorf("bad name: %s", state))
	}

	rs := make([]PartRange, 0)
	var r1, r2 PartRange
	r2 = start
	for _, rule := range wf.rules {
		r1, r2 = split(r2, rule)
		if r1.valid() {
			rs = append(rs, propagate(rule.target, r1, workflows)...)
		}
		if !r2.valid() {
			break
		}
	}
	if r2.valid() {
		rs = append(rs, propagate(wf.target, r2, workflows)...)
	}

	return rs
}

func main() {
	workflows := make(map[string]Workflow)
	parts := make([]Part, 0)
	re_workflow := regexp.MustCompile(`^(.+)\{(.*)\}$`)
	re_rule := regexp.MustCompile(`^(.)(<|>)(\d+):(.+)$`)
	re_part := regexp.MustCompile(`^\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}$`)

	helpers.ReadStdin(func(line string) {
		if m := re_workflow.FindStringSubmatch(line); m != nil {
			rulesArr := strings.Split(m[2], ",")
			fallback := rulesArr[len(rulesArr)-1]
			rulesArr = rulesArr[:len(rulesArr)-1]
			rules := make([]Rule, 0)
			for _, r := range rulesArr {
				mr := re_rule.FindStringSubmatch(r)
				if mr != nil {
					rules = append(rules, Rule{mr[1][0], mr[2] == "<", helpers.MustParseInt(mr[3]), mr[4]})
				}
			}
			workflows[m[1]] = Workflow{rules, fallback}
		} else if m := re_part.FindStringSubmatch(line); m != nil {
			parts = append(parts, Part{helpers.MustParseInt(m[1]), helpers.MustParseInt(m[2]), helpers.MustParseInt(m[3]), helpers.MustParseInt(m[4])})
		}
	})

	sum := int64(0)
	if helpers.Part1() {
		for _, part := range parts {
			if accepted(part, workflows) {
				sum += int64(part.x + part.m + part.a + part.s)
			}
		}
	} else {
		start := createPartRange(1, 4000)
		ranges := propagate("in", start, workflows)
		for _, r := range ranges {
			sum += r.count()
		}
	}
	fmt.Println(sum)
}
