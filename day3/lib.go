package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	batteryBanks := [][]int{}

	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatalf("Failed to open day 3 input: %v", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := []int{}
		line := scanner.Text()
		for char := range strings.SplitSeq(line, "") {
			digit, err := strconv.Atoi(char)
			if err != nil {
				log.Fatalf("Failed to read digit %s in line %s: %v", char, line, err)
			}

			bank = append(bank, digit)
		}

		batteryBanks = append(batteryBanks, bank)
	}

	return batteryBanks
}
