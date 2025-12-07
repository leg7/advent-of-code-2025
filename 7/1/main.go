package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"

	. "leg7.com/aoc2025/utils"
	"leg7.com/aoc2025/utils/assert"
)


func MatrixCoordCompareColThenRow(a, b MatrixCoord) int {
	if a.Col == b.Col {
		return cmp.Compare(a.Row, b.Row)
	} else if a.Col > b.Col {
		return 1
	} else {
		return -1
	}
}

func MatrixCoordCompareCol(a, b MatrixCoord) int {
	return cmp.Compare(a.Col, b.Col)
}

func MatrixCoordCompareRowThenCol(a, b MatrixCoord) int {
	if a.Row == b.Row {
		return cmp.Compare(a.Col, b.Col)
	} else if a.Row > b.Row {
		return 1
	} else {
		return -1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Find start symbol
	var startCoords MatrixCoord
	row := 0
	{
		regexpStart := regexp.MustCompile("S")
		assert.True(scanner.Scan(), "Could not read any input")
		firstLine := scanner.Bytes()

		loc := regexpStart.FindIndex(firstLine)
		assert.True(loc != nil, "Could not find start symbol S in the first line of the input")
		startCoords = MatrixCoord {
			Row: row,
			Col: loc[0],
		}
		assert.Equals(startCoords.Row, 0)
	}

	// Find splitters
	nodesByRow := []MatrixCoord{}
	for scanner.Scan() {
		row++

		for col, char := range scanner.Text() {
			if char == '^' {
				nodesByRow = append(nodesByRow, MatrixCoord {
					Row: row,
					Col: col,
				})
			}
		}
	}
	assert.True(len(nodesByRow) > 0, "Did not find any splitters in the input")

	// Sort them for a binary search
	nodesByCol := slices.Clone(nodesByRow)
	slices.SortFunc(nodesByCol, MatrixCoordCompareColThenRow)

	// Simulation
	// The problem is essentially a BFS walk through a binary DAG to count the number of nodes
	// ...
	// It turns out I didn't even need to do a BFS because the original nodes slices is sorted
	// by row and col so it corresponds to a topological walk which would have been simpler
	// I ended up making it simpler

	// Find root node
	position, _ := slices.BinarySearchFunc(nodesByCol, startCoords, MatrixCoordCompareColThenRow)
	rootNode := nodesByCol[position]
	assert.Equals(rootNode.Col, startCoords.Col)
	assert.True(rootNode.Row > startCoords.Row, "First splitter is not below the start symbol")

	reachable := map[MatrixCoord]bool{ rootNode: true }
	topologicalOrder := nodesByRow
	for _, current := range topologicalOrder {
		if !reachable[current] {
			continue
		}

		for col := current.Col - 1; col <= current.Col + 1; col += 2 {
			target := MatrixCoord { Row: current.Row, Col: col }
			position, found := slices.BinarySearchFunc(nodesByCol, target, MatrixCoordCompareColThenRow)
			assert.True(!found, "There cannot be two neighboring splitters in the input")

			neighborExists := position != len(nodesByCol) && nodesByCol[position].Col == target.Col
			if neighborExists {
				neighbor := nodesByCol[position]
				reachable[neighbor] = true
			}
		}
	}

	keys := make([]MatrixCoord, 0, len(reachable))
	for mc := range reachable {
		keys = append(keys, mc)
	}
	fmt.Println(len(keys))
}
