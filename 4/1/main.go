package main

import (
	"bufio"
	"os"
	"slices"
)

// Would be better to make these composable with a bitfield but there would be little benefit rn
// The order is weird to improve cache usage
type CardinalDirection uint8
const (
	West CardinalDirection = iota
	East
	NorthEast
	North
	NorthWest
	SouthEast
	South
	SouthWest
	CardinalDirectionCount
)

type Coord struct {
	col, row int
}

// Assumes that the origin of the grid is top left
var CardinalDirectionToOffset = []Coord {
	North: { 0, -1 },
	South: { 0, 1 },
	West: { -1, 0 },
	East: { 1, 0 },
	NorthWest: { -1, -1 },
	NorthEast: { 1, -1 },
	SouthWest: { -1, 1 },
	SouthEast: { 1, 1 },
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]byte{}
	for scanner.Scan() {
		// I was stuck on this all day because I didn't realise you have to clone scanner.Bytes()
		grid = append(grid, slices.Clone(scanner.Bytes()))
	}

	// This could be branchless if you process the extremeties of the grid
	// in separate loops but I don't have time rn
	rollsAccessible := 0
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == '.' {
				continue
			}

			neighborRolls := 0
			for direction := range CardinalDirectionCount {
				// Compute neighbor coords
				offset := CardinalDirectionToOffset[direction]
				neighborCoords := Coord {
					row: row + offset.row,
					col: col + offset.col,
				}

				rowOutOfBounds := neighborCoords.row < 0 || neighborCoords.row >= len(grid)
				colOutOfBounds := neighborCoords.col < 0 || neighborCoords.col >= len(grid[0])
				if rowOutOfBounds || colOutOfBounds {
					continue
				}

				neighborValue := grid[neighborCoords.row][neighborCoords.col]
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
