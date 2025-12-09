package day4

import "fmt"

func SolvePart2() {
	rolls := readInput()
	accessibleRollCount := 0

	foundAccessibleRoll := true

	for foundAccessibleRoll {
		foundAccessibleRoll = false
		for i, roll := range rolls {
			for j, char := range roll {
				if char == "@" && countAdjacentRolls(i, j, rolls) < 4 {
					accessibleRollCount++
					rolls[i][j] = "x"
					foundAccessibleRoll = true
				}
			}
		}
	}

	fmt.Println(accessibleRollCount)
}
