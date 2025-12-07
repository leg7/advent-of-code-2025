package utils

type MatrixCoord struct {
	Row, Col int
}

// Would be better to make these composable with a bitfield but there would be little benefit rn
// The order is weird to improve cache usage when walking through neighbors in a matrix
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

// Assumes that the origin of the grid is top left
var CardinalDirectionToOffset = []MatrixCoord {
	North: { 0, -1 },
	South: { 0, 1 },
	West: { -1, 0 },
	East: { 1, 0 },
	NorthWest: { -1, -1 },
	NorthEast: { 1, -1 },
	SouthWest: { -1, 1 },
	SouthEast: { 1, 1 },
}

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }

    if m == 1 {
        return n
    }

    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}


