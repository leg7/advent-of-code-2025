package main

import (
	"bufio"
	"os"

	. "leg7.com/aoc2025/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	batteryCount := 12

	maxJoltagesSum := 0
	for scanner.Scan() {
		joltages := scanner.Bytes()
		Assert(len(joltages) >= 2, "Input file must have atleast two or more batteries per bank but we only have %d in %s", len(joltages), string(joltages))

		value := 0
		start := 0
		pow := 100_000_000_000
		for neighborCount := batteryCount - 1; neighborCount >= 0; neighborCount-- {
			end := len(joltages) - neighborCount

			// Find the biggest candidate and it's index
			maxC, maxIdx := joltages[start], start
			for idx := start; idx < end; idx++ {
				if joltages[idx] > maxC {
					maxC, maxIdx = joltages[idx], idx
				}
			}

			start = maxIdx + 1
			value += pow * int(maxC - 48)
			pow /= 10
		}
		println(value)
		maxJoltagesSum += value
	}

	println(maxJoltagesSum)
}
