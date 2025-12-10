package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"leg7.com/aoc2025/utils/assert"
)

type Coord struct {
	X, Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	coords := []Coord{}

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		assert.Equals(len(fields), 2)

		x, err := strconv.Atoi(fields[0])
		assert.NoError(err)
		y, err := strconv.Atoi(fields[1])
		assert.NoError(err)

		coords = append(coords, Coord { x, y })
	}

	areaMax := 0
	for i := 0; i < len(coords) - 1; i++ {
		for j := i + 1; j < len(coords); j++ {
			area := (abs(coords[i].X - coords[j].X) + 1) * (abs(coords[i].Y - coords[j].Y) + 1)
			if area > areaMax {
				areaMax = area
			}
		}
	}

	fmt.Println(areaMax)
}
