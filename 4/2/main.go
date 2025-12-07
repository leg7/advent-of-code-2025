package main

import (
	"bufio"
	"os"
	"slices"
	"leg7.com/aoc2025/utils"
)

// Returns a pair (a, b) with a = 1 if the roll you called the function on was removed
// and b being the total ammount of neighbors removed
// This basically performs a BFS walk across the neighboring rolls and tries to remove them
func remove(grid [][]byte, row, col int, triedRemove map[utils.MatrixCoord]bool) (int, int) {
	if grid[row][col] != '@' { // in case a roll in the queue was changed to 'x' by another recursive call
		return 0, 0
	}

	current := utils.MatrixCoord{ Row: row, Col: col }
	triedRemove[current] = true

	neighborRollsCount := 0
	neighborRollsToTryToRemove := []utils.MatrixCoord{}
	for direction := range utils.CardinalDirectionCount {
		offset := utils.CardinalDirectionToOffset[direction]
		neighborCoords := utils.MatrixCoord {
			Row: current.Row + offset.Row,
			Col: current.Col + offset.Col,
		}

		rowOutOfBounds := neighborCoords.Row < 0 || neighborCoords.Row >= len(grid)
		colOutOfBounds := neighborCoords.Col < 0 || neighborCoords.Col >= len(grid[0])
		if rowOutOfBounds || colOutOfBounds {
			continue
		}

		neighborValue := grid[neighborCoords.Row][neighborCoords.Col]
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
			removed, removedSubtotal := remove(grid, neighbor.Row, neighbor.Col, triedRemove)
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
				triedRemove := map[utils.MatrixCoord]bool{}
				_, rollsRemoved := remove(grid, row, col, triedRemove)
				rollsRemovedTotal += rollsRemoved
			}
		}
	}

	println(rollsRemovedTotal)
}
