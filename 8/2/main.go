package main

import (
	"bufio"
	"cmp"
	"fmt"
	"iter"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"leg7.com/aoc2025/utils/assert"
)

type Vec3 struct {
	X, Y, Z float64
}

func DistanceSquared(a, b Vec3) float64 {
	return math.Pow(a.X - b.X, 2) + math.Pow(a.Y - b.Y, 2) + math.Pow(a.Z - b.Z, 2)
}

type Segment struct {
	A, B Vec3
}

func (s *Segment) Length() float64 {
	return DistanceSquared(s.A, s.B)
}

type Graph[T comparable] struct {
	AdjacencyList map[T][]T
}

func NewGraph[T comparable]() Graph[T] {
	return Graph[T] {
		AdjacencyList: map[T][]T{},
	}
}

func (g *Graph[T]) AddEdge(from, to T) {
	g.AdjacencyList[from] = append(g.AdjacencyList[from], to)
	g.AdjacencyList[to]   = append(g.AdjacencyList[to], from)
}

func (g *Graph[T]) AddNode(node T) {
	g.AdjacencyList[node] = []T{}
}

func Bfs[T comparable](source T, adjacencyList map[T][]T) iter.Seq[T] {
	return func (yield func(T) bool) {
		queue := adjacencyList[source]
		visited := map[T]bool{}
		for _, node := range adjacencyList[source] {
			visited[node] = true
		}

		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			if !yield(current) {
				return
			}

			for _, neighbor := range adjacencyList[current] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
	}
}

func (g *Graph[T]) Reachables(from T) iter.Seq[T] {
	return Bfs(from, g.AdjacencyList)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read points in input
	points := []Vec3{}

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		assert.Equals(len(coords), 3)

		x, err := strconv.ParseFloat(coords[0], 32)
		assert.NoError(err)
		y, err := strconv.ParseFloat(coords[1], 32)
		assert.NoError(err)
		z, err := strconv.ParseFloat(coords[2], 32)
		assert.NoError(err)

		points = append(points, Vec3{ x, y, z })
	}
	assert.True(len(points) > 0, "We could not read any points")

	// Calculate all possible segments
	segments := []Segment{}

	for i := 0; i < len(points) - 1; i++ {
		for j := i + 1; j < len(points); j++ {
			segments = append(segments, Segment { points[i], points[j] })
		}
	}

	slices.SortFunc(segments, func (a, b Segment) int {
		return cmp.Compare(a.Length(), b.Length())
	})

	// Build circuit graph

	circuitGraph := NewGraph[Vec3]()

	for _, segment := range segments {
		circuitGraph.AddEdge(segment.A, segment.B)

		circuitSizeMax := 0
		visited := map[Vec3]bool{}

		for _, point := range points {
			if visited[point] {
				continue
			}
			visited[point] = true

			circuitSize := 0

			for neighbor := range circuitGraph.Reachables(point) {
				visited[neighbor] = true
				circuitSize++
			}

			circuitSizeMax = max(circuitSizeMax, circuitSize)
		}

		if circuitSizeMax == len(points) {
			fmt.Println(int(segment.A.X) * int(segment.B.X))
			break
		}
	}
}
