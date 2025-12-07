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
	// Part 2 asks how many different branches the DAG has
	// The examples show a DFS walk over all branches and while it works for the small example
	// It would take forever on a deep DAG like the input the growth is exponentially
	// with every new level so that was just bait
	//
	// I was trying to find a formula to calculate the number of branches with other properties
	// But I spoiled myself on reddit and it turns out we can just use pascal's triangle to get
	// the answer since this is a binary DAG
	//
	// However if splitters were not binary then it would be some sort of extended's pascal triangle
	// This is apparently a common DP problem equivalent to galton boards and lattice paths

	// Find root node
	position, _ := slices.BinarySearchFunc(nodesByCol, startCoords, MatrixCoordCompareColThenRow)
	rootNode := nodesByCol[position]
	assert.Equals(rootNode.Col, startCoords.Col)
	assert.True(rootNode.Row > startCoords.Row, "First splitter is not below the start symbol")

	pathsToNode := map[MatrixCoord]int{ rootNode: 1 }
	paths := 0
	topologicalOrder := nodesByRow
	for _, current := range topologicalOrder {
		currentIsUnreachable := pathsToNode[current] == 0
		if currentIsUnreachable {
			continue
		}

		closedBranches := 2
		for col := current.Col - 1; col <= current.Col + 1; col += 2 {
			target := MatrixCoord { Row: current.Row, Col: col }
			position, found := slices.BinarySearchFunc(nodesByCol, target, MatrixCoordCompareColThenRow)
			assert.True(!found, "There cannot be two neighboring splitters in the input")

			neighborExists := position != len(nodesByCol) && nodesByCol[position].Col == target.Col
			neighborBelow := neighborExists && nodesByCol[position].Row > current.Row
			if neighborBelow {
				neighbor := nodesByCol[position]
				pathsToNode[neighbor] += pathsToNode[current]
				closedBranches--
			}
		}

		paths += pathsToNode[current] * closedBranches
	}

	fmt.Println(paths)
}
