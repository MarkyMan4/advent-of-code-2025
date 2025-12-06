package day2

import (
	"fmt"
	"strconv"
)

func SolvePart1() {
	ranges := readInput()
	sum := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 != 0 {
				continue
			}

			mid := len(id) / 2
			if id[:mid] == id[mid:] {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}
