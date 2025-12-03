package day1

import (
	"fmt"
	"math"
)

func SolvePart2() {
	currentDialPos := dialStart
	instructions := readInput()
	zeroCount := 0 // number of times we landed on zero

	for _, inst := range instructions {
		// if instruction is L256, we know we've already passed zero twice
		zeroCount += int(math.Floor(float64(inst.clicks) / 100))

		// mod 100 to get the number to move
		// if number is 988, moving 900 lands us back at the same spot
		// so we only care about the 88
		clicks := inst.clicks % 100

		switch inst.direction {
		case 'R':
			currentDialPos += clicks

			if currentDialPos > maxDialPos {
				currentDialPos = (currentDialPos % maxDialPos) - 1
				zeroCount++
			}
		case 'L':
			if clicks > currentDialPos {
				currentDialPos = maxDialPos - (clicks - currentDialPos - 1)
				zeroCount++
			} else {
				currentDialPos -= clicks
			}

			if currentDialPos > maxDialPos {
				currentDialPos = (currentDialPos % maxDialPos) - 1
			}
		}
	}

	fmt.Println(zeroCount)
}
