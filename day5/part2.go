package day5

import (
	"cmp"
	"fmt"
	"slices"
)

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func SolvePart2() {
	ranges, _ := readInput()

	// sort ranges by lower bound
	slices.SortFunc(ranges, func(r1, r2 *idRange) int {
		return cmp.Compare(r1.lowerBound, r2.lowerBound)
	})

	// some ranges may have overlap, combine them to make the next step of finding all the number in a range easier
	combinedRanges := []*idRange{}
	for i := 0; i < len(ranges); i++ {
		curRange := ranges[i]
		for i < len(ranges)-1 && curRange.upperBound >= ranges[i+1].lowerBound {
			curRange.upperBound = maxInt(curRange.upperBound, ranges[i+1].upperBound)
			i++
		}

		combinedRanges = append(combinedRanges, curRange)
	}

	freshCount := 0
	for _, r := range combinedRanges {
		freshCount += (r.upperBound - r.lowerBound) + 1
	}

	fmt.Println(freshCount)
}
