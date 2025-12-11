package day7

import "fmt"

func SolvePart1() {
	grid := readInput()
	splitCount := 0

	for i, line := range grid {
		if i == 0 {
			continue
		}

		splitsThisLine := 0
		for j, cell := range line {
			above := grid[i-1][j]

			if cell == SPLITTER && above == BEAM {
				splitCount++
				splitsThisLine++
				grid[i][j-1] = BEAM
				grid[i][j+1] = BEAM

			} else if (above == START || above == BEAM) && cell == EMPTY {
				grid[i][j] = BEAM
			}
		}
	}

	for _, line := range grid {
		fmt.Println(line)
	}

	fmt.Println(splitCount)
}
