package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func readInput() []idRange {
	ranges := []idRange{}

	content, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		log.Fatalf("Unable to read day 2 input %v", err)
	}

	for _, rangeStr := range strings.Split(string(content), ",") {
		bounds := strings.Split(rangeStr, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatalf("Failed to parse start of range %s: %v", rangeStr, err)
		}

		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatalf("Failed to parse end of range %s: %v", rangeStr, err)
		}

		ranges = append(ranges, idRange{start, end})
	}

	return ranges
}
