package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"leg7.com/aoc2025/utils/assert"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxJoltagesSum := 0
	for scanner.Scan() {
		joltages := scanner.Bytes()
		assert.True(len(joltages) >= 2, "Input file must have atleast two or more batteries per bank but we only have %d in %s", len(joltages), string(joltages))

		outer:
		for dozens := 9; dozens >= 1; dozens-- {
			for units := 9; units >= 0; units-- {
				regex := fmt.Sprintf("%d.*%d", dozens, units)
				matched, _ := regexp.Match(regex, joltages)

				if matched {
					maxJoltage := dozens * 10 + units
					maxJoltagesSum += maxJoltage
					break outer
				}
			}
		}
	}

	println(maxJoltagesSum)
}
