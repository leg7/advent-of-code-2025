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

// Returns a pair (a, b) with a = 1 if the roll you called the function on was removed
// and b being the total ammount of neighbors removed
// This basically performs a BFS walk across the neighboring rolls and tries to remove them
func remove(grid [][]byte, row, col int, triedRemove map[Coord]bool) (int, int) {
	if grid[row][col] != '@' { // in case a roll in the queue was changed to 'x' by another recursive call
		return 0, 0
	}

	current := Coord{ row: row, col: col }
	triedRemove[current] = true

	neighborRollsCount := 0
	neighborRollsToTryToRemove := []Coord{}
	for direction := range CardinalDirectionCount {
		offset := CardinalDirectionToOffset[direction]
		neighborCoords := Coord {
			row: current.row + offset.row,
			col: current.col + offset.col,
		}

		rowOutOfBounds := neighborCoords.row < 0 || neighborCoords.row >= len(grid)
		colOutOfBounds := neighborCoords.col < 0 || neighborCoords.col >= len(grid[0])
		if rowOutOfBounds || colOutOfBounds {
			continue
		}

		neighborValue := grid[neighborCoords.row][neighborCoords.col]
		neighborIsARoll := neighborValue == '@'
		if neighborIsARoll {
			neighborRollsCount++

			if !triedRemove[neighborCoords] {
				neighborRollsToTryToRemove = append(neighborRollsToTryToRemove, neighborCoords)
			}
		}
	}

	if neighborRollsCount < 4 {
		grid[row][col] = 'x'
		return 1, 1
	} else {
		removedDirectNeighbors := 0
		removedTotal := 0
		for _, neighbor := range neighborRollsToTryToRemove {
			removed, removedSubtotal := remove(grid, neighbor.row, neighbor.col, triedRemove)
			removedTotal += removedSubtotal
			removedDirectNeighbors += removed
		}

		if neighborRollsCount - removedDirectNeighbors < 4 {
			grid[row][col] = 'x'
			removedTotal++
			return 1, removedTotal
		} else {
			return 0, removedTotal
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]byte{}
	for scanner.Scan() {
		grid = append(grid, slices.Clone(scanner.Bytes()))
	}

	rollsRemovedTotal := 0
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == '@' {
				triedRemove := map[Coord]bool{}
				_, rollsRemoved := remove(grid, row, col, triedRemove)
				rollsRemovedTotal += rollsRemoved
			}
		}
	}

	println(rollsRemovedTotal)
}
