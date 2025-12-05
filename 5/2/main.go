package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"leg7.com/aoc2025/utils/assert"
)

type Range struct {
	lb, ub int64
}

func RangeCmp(a, b Range) int {
	return cmp.Compare(a.lb, b.lb)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ranges := []Range{}

	// Read ranges
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		rangeStr := strings.Split(scanner.Text(), "-")
		assert.Equals(len(rangeStr), 2)

		lb, err := strconv.ParseInt(rangeStr[0], 10, 64)
		assert.NoError(err)

		ub, err := strconv.ParseInt(rangeStr[1], 10, 64)
		assert.NoError(err)

		ranges = append(ranges, Range { lb: lb, ub: ub })
	}
	assert.True(len(ranges) > 0, "No ranges have been read\n")

	// Merge overlapping ranges
	slices.SortFunc(ranges, RangeCmp)
	rangesCompressed := []Range{ ranges[0] }

	for idx := 1; idx < len(ranges); idx++ {
		current := &rangesCompressed[len(rangesCompressed) - 1]
		candidate := &ranges[idx]

		overlaps := candidate.lb >= current.lb && candidate.lb <= current.ub
		proceeds := candidate.lb == current.ub + 1
		if overlaps || proceeds {
			current.ub = max(candidate.ub, current.ub)
		} else {
			rangesCompressed = append(rangesCompressed, *candidate)
		}
	}

	assert.True(slices.IsSortedFunc(rangesCompressed, RangeCmp), "Compressed ranges slice is not ordered")

	fmt.Printf("Went from %d ranges to %d compressed ranges\n", len(ranges), len(rangesCompressed))

	// Read ingredients
	totalRangesSpan := int64(0)
	for _, r := range rangesCompressed {
		totalRangesSpan += r.ub - r.lb + 1
	}

	println(totalRangesSpan)
}
