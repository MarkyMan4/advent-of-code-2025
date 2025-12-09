package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MarkyMan4/advent-of-code-2025/day1"
	"github.com/MarkyMan4/advent-of-code-2025/day2"
	"github.com/MarkyMan4/advent-of-code-2025/day3"
	"github.com/MarkyMan4/advent-of-code-2025/day4"
	"github.com/MarkyMan4/advent-of-code-2025/day5"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("provide required positional arguments: <day> <part>")
		os.Exit(1)
	}

	args := os.Args

	day, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("day must be an integer")
		os.Exit(1)
	}

	part, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("part must be an integer")
		os.Exit(1)
	}

	solveDayAndPart(day, part)
}

func solveDayAndPart(day int, part int) {
	switch day {
	case 1:
		if part == 1 {
			day1.SolvePart1()
		} else if part == 2 {
			day1.SolvePart2()
		}
	case 2:
		if part == 1 {
			day2.SolvePart1()
		} else if part == 2 {
			day2.SolvePart2()
		}
	case 3:
		if part == 1 {
			day3.SolvePart1()
		} else if part == 2 {
			day3.SolvePart2()
		}
	case 4:
		if part == 1 {
			day4.SolvePart1()
		} else if part == 2 {
			day4.SolvePart2()
		}
	case 5:
		if part == 1 {
			day5.SolvePart1()
		} else if part == 2 {
			day5.SolvePart2()
		}
	}
}
