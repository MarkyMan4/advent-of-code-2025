package day7

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	START    = "S"
	EMPTY    = "."
	BEAM     = "|"
	SPLITTER = "^"
)

func readInput() [][]string {
	lines := [][]string{}
	file, err := os.Open("inputs/day7.txt")
	if err != nil {
		log.Fatalf("Failed to read day 7 input: %v", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines
}
