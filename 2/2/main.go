package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	. "leg7.com/aoc2025/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	splitFunc := func(data []byte, atEof bool) (advance int, token []byte, err error) {
		if atEof && len(data) == 0 {
			return 0, nil, bufio.ErrFinalToken
		}

		token, advance, err = []byte{}, 0, nil
		foundSep := false
		for ; !foundSep && advance < len(data); advance++ {
			if data[advance] == '-' || data[advance] == ',' {
				foundSep = true
			} else if data[advance] != '\n' && data[advance] != '\r' {
				token = append(token, data[advance])
			}
		}

		if !foundSep && atEof {
			err = bufio.ErrFinalToken
		}

		return
	}
	scanner.Split(splitFunc)

	invalidIdsSum := 0
	for scanner.Scan() {
		lb, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ub, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		Assert(lb <= ub, "Lower bound %d > Upper bound %d\n", lb, ub)

		for id := lb; id <= ub; id++ {
			iStr := strconv.Itoa(id)
			stream := iStr + iStr
			stream = stream[1:len(stream)-1]
			if strings.Contains(stream, iStr) {
				println("Invalid id : ", id)
				invalidIdsSum += id
			}
		}
	}
	println(invalidIdsSum)
}
