package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MarkyMan4/advent-of-code-2025/day1"
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
		switch part {
		case 1:
			day1.SolvePart1()
		case 2:
			day1.SolvePart2()
		}
	}
}
