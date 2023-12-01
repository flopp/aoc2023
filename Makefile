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
	@echo "expected: ?"
	$(run1)
	@echo "expected: ?"
	$(run2)