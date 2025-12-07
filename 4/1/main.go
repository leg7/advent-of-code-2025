package main

import (
	"bufio"
	"os"
	"slices"
	"leg7.com/aoc2025/utils"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]byte{}
	for scanner.Scan() {
		// I was stuck on this all day because I didn't realise you have to clone scanner.Bytes()
		grid = append(grid, slices.Clone(scanner.Bytes()))
	}

	// This could have less branches if you process the extremeties of the grid
	// in separate loops but I don't have time rn
	rollsAccessible := 0
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == '.' {
				continue
			}

			neighborRolls := 0
			for direction := range utils.CardinalDirectionCount {
				// Compute neighbor coords
				offset := utils.CardinalDirectionToOffset[direction]
				neighborCoords := utils.MatrixCoord {
					Row: row + offset.Row,
					Col: col + offset.Col,
				}

				rowOutOfBounds := neighborCoords.Row < 0 || neighborCoords.Row >= len(grid)
				colOutOfBounds := neighborCoords.Col < 0 || neighborCoords.Col >= len(grid[0])
				if rowOutOfBounds || colOutOfBounds {
					continue
				}

				neighborValue := grid[neighborCoords.Row][neighborCoords.Col]
				neighborIsARoll := neighborValue == '@' || neighborValue == 'x'
				if neighborIsARoll {
					neighborRolls++
				}
			}

			if neighborRolls < 4 {
				rollsAccessible++
				grid[row][col] = 'x'
			}
		}
	}

	println(rollsAccessible)
	for _, l := range grid {
		println(string(l))
	}
}
