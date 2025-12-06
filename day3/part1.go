package day3

import (
	"fmt"
	"slices"
	"strconv"
)

func SolvePart1() {
	banks := readInput()
	totalJoltage := 0

	for _, bank := range banks {
		highest := slices.Max(bank[:len(bank)-1])
		highestIdx := slices.Index(bank, highest)
		nextHighest := slices.Max(bank[highestIdx+1:])
		joltage, _ := strconv.Atoi(strconv.Itoa(highest) + strconv.Itoa(nextHighest))

		totalJoltage += joltage
	}

	fmt.Println(totalJoltage)
}
