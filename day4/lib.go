package day4

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readInput() [][]string {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		log.Fatalf("Failed to read day 4 input: %v", err)
	}

	lines := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines
}

func countAdjacentRolls(i int, j int, rolls [][]string) int {
	// count how many rolls are in the 8 cells adjacent to i, j
	count := 0

	// above
	if i > 0 && rolls[i-1][j] == "@" {
		count++
	}

	// above right
	if i > 0 && j < len(rolls[i])-1 && rolls[i-1][j+1] == "@" {
		count++
	}

	// right
	if j < len(rolls[i])-1 && rolls[i][j+1] == "@" {
		count++
	}

	// down right
	if i < len(rolls)-1 && j < len(rolls[i])-1 && rolls[i+1][j+1] == "@" {
		count++
	}

	// down
	if i < len(rolls)-1 && rolls[i+1][j] == "@" {
		count++
	}

	// down left
	if i < len(rolls)-1 && j > 0 && rolls[i+1][j-1] == "@" {
		count++
	}

	// left
	if j > 0 && rolls[i][j-1] == "@" {
		count++
	}

	// above left
	if i > 0 && j > 0 && rolls[i-1][j-1] == "@" {
		count++
	}

	return count
}
