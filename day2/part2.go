package day2

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func isIdValid(id int) bool {
	idStr := strconv.Itoa(id)
	seq := ""

	for i, char := range idStr {
		// if we made it past first half of the id, it must be valid
		if i >= int(math.Ceil(float64(len(idStr))/2)) {
			return true
		}

		seq += string(char)

		// check if the length of the current substring is divisible
		// by the length of the id so we know if it's possible if the
		// id could be made up of this sequence
		if len(idStr)%len(seq) != 0 {
			continue
		}

		// check if id is composed of repetitions of the current substring
		re := regexp.MustCompile(fmt.Sprintf("^(%s){2,}$", seq))
		if re.MatchString(idStr) {
			return false
		}
	}

	return true
}

func SolvePart2() {
	ranges := readInput()
	sum := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if !isIdValid(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}
