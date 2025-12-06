package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// FYI this same approach could be used to solve part 1, or any arbitrary number of batteries needed
// assuming BATTERIES_REQUIRED < len(bank)

const BATTERIES_REQUIRED = 12

func SolvePart2() {
	// find twelve digits that form the biggest number (must be read left to right)
	// find the highest number that is at least 12 spaces from the end
	// then the highest number that is at least 11 spaces from the end
	// repeat until we have 12 digits
	banks := readInput()
	totalJoltage := 0

	for _, bank := range banks {
		digits := []string{}
		minIdx := 0

		for len(digits) < BATTERIES_REQUIRED {
			remaining := BATTERIES_REQUIRED - len(digits) - 1
			highest := slices.Max(bank[minIdx : len(bank)-remaining])
			digits = append(digits, strconv.Itoa(highest))
			minIdx += slices.Index(bank[minIdx:], highest) + 1
		}

		joltage, _ := strconv.Atoi(strings.Join(digits, ""))
		totalJoltage += joltage
	}

	fmt.Println(totalJoltage)
}
