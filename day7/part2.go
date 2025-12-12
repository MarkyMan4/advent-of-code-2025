package day7

import "fmt"

type gridPos struct {
	i int
	j int
}

func countTimelines(grid [][]string, startPos gridPos, memo map[gridPos]int) int {
	for startPos.i < len(grid) && grid[startPos.i][startPos.j] != SPLITTER {
		startPos.i++
	}

	if startPos.i >= len(grid)-1 {
		return 1
	}

	if _, ok := memo[startPos]; !ok {
		memo[startPos] = countTimelines(grid, gridPos{startPos.i, startPos.j - 1}, memo) +
			countTimelines(grid, gridPos{startPos.i, startPos.j + 1}, memo)
	}

	return memo[startPos]
}

func SolvePart2() {
	grid := readInput()
	startPos := gridPos{0, 0}

	// find position of S
	for j := range grid[0] {
		if grid[0][j] == START {
			startPos.j = j
		}
	}

	var memo = make(map[gridPos]int)
	fmt.Println(countTimelines(grid, startPos, memo))
}
