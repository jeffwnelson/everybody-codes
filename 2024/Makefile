## When nothing is passed, default to help
.DEFAULT_GOAL := help

## Detect year from current directory name
YEAR := $(notdir $(CURDIR))

## Force to run even if conditions are true
FORCE:

## Show available commands
help:
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@echo "  quest<XX>      Run solution for quest XX (e.g., make quest01)"
	@echo "  new<XX>        Create new quest XX with template (e.g., make new01)"
	@echo "  delete<XX>     Delete quest XX and its inputs (e.g., make delete01)"
	@echo "  help           Show this help message"

## Run a quest: make quest01
quest%: FORCE
	@go run ./quest$*

## Create a new quest: make new01
new%: FORCE
	@mkdir -p quest$* inputs/quest$*
	@touch inputs/quest$*/part1.txt inputs/quest$*/part2.txt inputs/quest$*/part3.txt
	@if [ ! -f quest$*/solution.go ]; then \
		echo 'package main\n\nimport (\n\t"everybody-codes/utils"\n)\n\nconst (\n\tinputPart1 = "$(YEAR)/inputs/quest$*/part1.txt"\n\tinputPart2 = "$(YEAR)/inputs/quest$*/part2.txt"\n\tinputPart3 = "$(YEAR)/inputs/quest$*/part3.txt"\n)\n\nfunc main() {\n\tutils.Run("Part 1", part1)\n\tutils.Run("Part 2", part2)\n\tutils.Run("Part 3", part3)\n}\n\nfunc part1() int {\n\treturn 0\n}\n\nfunc part2() int {\n\treturn 0\n}\n\nfunc part3() int {\n\treturn 0\n}' > quest$*/solution.go; \
	fi
	@echo "Created quest$*/ and inputs/quest$*/"

## Delete a quest: make delete01
delete%: FORCE
	@read -p "Delete quest$* and inputs/quest$*? [y/N] " confirm && [ "$$confirm" = "y" ] && rm -rf quest$* inputs/quest$* && echo "Deleted quest$*/ and inputs/quest$*/" || echo "Cancelled"
