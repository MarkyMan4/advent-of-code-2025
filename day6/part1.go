package day6

import "fmt"

func SolvePart1() {
	operands, operators := readInput()
	total := 0
	col := 0

	for col < len(operands[0]) {
		nums := []int{}
		for i := 0; i < len(operands); i++ {
			nums = append(nums, operands[i][col])
		}

		total += reduce(nums, operators[col])
		col++
	}

	fmt.Println(total)
}
