package day6

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func transpose(arr []string) []string {
	newArr := []string{}
	for i, s := range arr {
		for j, char := range s {
			if i == 0 {
				newArr = append(newArr, string(char))
			} else {
				newArr[j] += string(char)
			}
		}
	}

	return newArr
}

func SolvePart2() {
	total := 0
	grid := readInputPart2()
	numberGrid := grid[:len(grid)-1]
	operators := strings.Fields(grid[len(grid)-1])

	numberGrid = transpose(numberGrid)
	opIdx := 0
	operands := []int{}

	for i, line := range numberGrid {
		maybeNum := strings.TrimSpace(line)

		if len(maybeNum) > 0 {
			num, err := strconv.Atoi(maybeNum)
			if err != nil {
				log.Fatalf("Failed to parse int from %s: %v", maybeNum, err)
			}

			operands = append(operands, num)
		}

		if len(maybeNum) == 0 || i == len(numberGrid)-1 {
			// if we hit an empty line or the end of the numbers, apply the operaator and add to total
			result := reduce(operands, operators[opIdx])
			total += result
			opIdx++
			operands = []int{}
			continue
		}
	}

	fmt.Println(total)
}
