define run_test
	@go run day$@/main.go part$(1) test < day$@/$(2).txt
endef
define run_puzzle
	@go run day$@/main.go part$(1) puzzle < day$@/$(2).txt
endef
define run1_puzzle
	@echo "=>"
	$(call run_puzzle,1,puzzle)
	@echo
endef
define run2_puzzle
	@echo "=>"
	$(call run_puzzle,2,puzzle)
	@echo
endef
	
define run1
	$(call run_test,1,test)
	$(run1_puzzle)
endef

define run2
	$(call run_test,2,test)
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
	$(call run_test,2,test2)
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
	$(call run_test,1,test)
	@echo "expected: 6"
	$(call run_test,1,test2)
	$(run1_puzzle)
	@echo "expected: 6"
	$(call run_test,2,test3)
	$(run2_puzzle)

.PHONY: 09
09:
	@echo "expected: 114"
	$(run1)
	@echo "expected: 2"
	$(run2)

.PHONY: 10
10:
	@echo "expected: 4"
	$(call run_test,1,test)
	@echo "expected: 8"
	$(call run_test,1,test2)
	$(run1_puzzle)

	@echo "expected: 4"
	$(call run_test,2,test2a)
	@echo "expected: 8"
	$(call run_test,2,test2b)
	@echo "expected: 10"
	$(call run_test,2,test2c)
	$(run2_puzzle)

.PHONY: 11
11:
	@echo "expected: 374"
	$(run1)
	@echo "expected: 1030"
	$(run2)

.PHONY: 12
12:
	@echo "expected: 21"
	$(run1)
	@echo "expected: 525152"
	$(run2)
	
	

.PHONY: 13
13:
	@echo "expected: 405 / 4006"
	$(run1)
	@echo "expected: 400 / 28627"
	$(run2)
