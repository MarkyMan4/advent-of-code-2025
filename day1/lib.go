package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dialStart int = 50
const minDialPos int = 0
const maxDialPos int = 99

type dialRotation struct {
	direction rune
	clicks    int
}

func readInput() []dialRotation {
	content, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(1)
	}

	instructions := []dialRotation{}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		direction := rune(line[0])
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("failed to parse number from line")
			os.Exit(1)
		}

		instructions = append(instructions, dialRotation{direction, clicks})
	}

	return instructions
}
