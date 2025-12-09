package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	lowerBound int
	upperBound int
}

// returns a slice of id ranges and a slice of IDs
func readInput() ([]*idRange, []int) {
	file, err := os.Open("inputs/day5.txt")
	if err != nil {
		log.Fatalf("Failed to read day 5 input: %v", err)
	}

	scanner := bufio.NewScanner(file)
	// read id ranges
	idRanges := []*idRange{}
	for scanner.Scan() {
		line := scanner.Text()

		// if we reach a blank line, we've read all id ranges
		if len(line) == 0 {
			break
		}

		bounds := strings.Split(line, "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatalf("Failed to convert lower bound to int in range %s: %v", bounds, err)
		}

		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatalf("Failed to convert upper bound to int in range %s: %v", line, err)
		}

		idRanges = append(idRanges, &idRange{lower, upper})
	}

	// read ingredientIds
	ingredientIds := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to id to int %s: %v", line, err)
		}

		ingredientIds = append(ingredientIds, id)
	}

	return idRanges, ingredientIds
}
