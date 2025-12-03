package day1

import (
	"fmt"
)

func SolvePart1() {
	currentDialPos := dialStart
	instructions := readInput()
	zeroCount := 0 // number of times we landed on zero

	for _, inst := range instructions {
		// mod 100 to get the number to move
		// if number is 988, moving 900 lands us back at the same spot
		// so we only care about the 88
		clicks := inst.clicks % 100

		switch inst.direction {
		case 'R':
			currentDialPos += clicks

			if currentDialPos > maxDialPos {
				currentDialPos = (currentDialPos % maxDialPos) - 1
			}
		case 'L':
			if clicks > currentDialPos {
				currentDialPos = maxDialPos - (clicks - currentDialPos - 1)
			} else {
				currentDialPos -= clicks
			}

			if currentDialPos < minDialPos {
				currentDialPos = (currentDialPos % maxDialPos) + 1
			}
		}

		if currentDialPos == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
