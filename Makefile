define run1_test
	@go run day$@/main.go part1 test < day$@/test.txt
endef
define run1_test2
	@go run day$@/main.go part1 test < day$@/test2.txt
endef
define run1_puzzle
	@echo "=>"
	@go run day$@/main.go part1 puzzlr < day$@/puzzle.txt
	@echo
endef

define run2_test
	@go run day$@/main.go part2 test < day$@/test.txt
endef
define run2_test2
	@go run day$@/main.go part2 test < day$@/test2.txt
endef
define run2_test3
	@go run day$@/main.go part2 test < day$@/test3.txt
endef
define run2_puzzle
	@echo "=>"
	@go run day$@/main.go part2 puzzlr < day$@/puzzle.txt
	@echo
endef
	
define run1
	$(run1_test)
	$(run1_puzzle)
endef

define run2
	$(run2_test)
	$(run2_puzzle)
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
	$(run2_puzzle)

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

.PHONY: 07
07:
	@echo "expected: 6440"
	$(run1)
	@echo "expected: 5905"
	$(run2)

.PHONY: 08
08:
	@echo "expected: 2"
	$(run1_test)
	@echo "expected: 6"
	$(run1_test2)
	$(run1_puzzle)
	@echo "expected: 6"
	$(run2_test3)
	$(run2_puzzle)
