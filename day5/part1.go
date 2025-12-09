package day5

import "fmt"

func isIngredientFresh(ranges []*idRange, id int) bool {
	// check if the id is contained in any of the id ranges
	for _, idRange := range ranges {
		if id >= idRange.lowerBound && id <= idRange.upperBound {
			return true
		}
	}

	return false
}

func SolvePart1() {
	ranges, ids := readInput()
	freshCount := 0

	for _, id := range ids {
		if isIngredientFresh(ranges, id) {
			freshCount++
		}
	}

	fmt.Println(freshCount)
}
