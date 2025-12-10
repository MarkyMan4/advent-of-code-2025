package day6

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func reduce(nums []int, op string) int {
	acc := nums[0]
	for i := 1; i < len(nums); i++ {
		if op == "*" {
			acc *= nums[i]
		} else if op == "+" {
			acc += nums[i]
		}
	}

	return acc
}

func readInput() ([][]int, []string) {
	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		log.Fatalf("Failed to read day 6 input: %v", err)
	}

	scanner := bufio.NewScanner(file)
	operands := [][]int{}
	operators := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if values[0] == "+" || values[0] == "*" {
			operators = values
			continue
		}

		operandLine := []int{}
		for _, val := range values {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("Failed to convert value %s to int: %v", val, err)
			}

			operandLine = append(operandLine, num)
		}

		operands = append(operands, operandLine)
	}

	return operands, operators
}

func readInputPart2() []string {
	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		log.Fatalf("Failed to read day 6 input: %v", err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
