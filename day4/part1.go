package day4

import "fmt"

func SolvePart1() {
	rolls := readInput()
	accessibleRollCount := 0

	for i, roll := range rolls {
		for j, cell := range roll {
			if cell == "@" && countAdjacentRolls(i, j, rolls) < 4 {
				accessibleRollCount++
			}
		}
	}

	fmt.Println(accessibleRollCount)
}
