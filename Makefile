define run1
	@go run day$@/main.go part1 test < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part1 puzzle < day$@/puzzle.txt
	@echo
endef

define run2
	@go run day$@/main.go part2 test < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part2 puzzle < day$@/puzzle.txt
	@echo
endef

define run2_test2
	@go run day$@/main.go part2 test < day$@/test2.txt
	@echo "=>"
	@go run day$@/main.go part2 puzzle < day$@/puzzle.txt
	@echo
endef

all:
	@echo "Run 'make dayXX' to create a template directory for day XX"
	@echo "Run 'make XX' to run the test and puzzle inputs on the day XX solution"

day%:
	@go run tools/create_day/main.go $@

.PHONY: 01
01:
	@echo "expected: 142"
	$(run1)
	@echo "expected: 281"
	$(run2_test2)

.PHONY: 02
02:
	@echo "expected: 8"
	$(run1)
	@echo "expected: 2286"
	$(run2)


.PHONY: 03
03:
	@echo "expected: 4361"
	$(run1)
	@echo "expected: 467835"
	$(run2)

.PHONY: 04
04:
	@echo "expected: 13"
	$(run1)
	@echo "expected: 30"
	$(run2)

.PHONY: 05
05:
	@echo "expected: 35"
	$(run1)
	@echo "expected: 46"
	$(run2)

.PHONY: 06
06:
	@echo "expected: 288"
	$(run1)
	@echo "expected: 71503"
	$(run2)
