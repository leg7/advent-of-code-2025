package main

import (
	"bufio"
	"os"
	"strconv"
	. "leg7.com/aoc2025/utils"
)

func odd(n int) bool {
	return n % 2 != 0
}

func digitLength(n int) int {
	dividend := 10
	length := 1
	for n / dividend != 0 {
		length++
		dividend *= 10
	}
	return length
}

func roundUpToEvenLength(n int) int {
	length := digitLength(n)
	if odd(length) {
		n += IntPow(10, length) - n
	}
	Assert(!odd(digitLength(n)), "Digit length of %d cannot be odd after rounding it up to even length", n)
	return n
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

		for id := roundUpToEvenLength(lb); id <= ub; {
			digitLen := digitLength(id)
			iStr := strconv.Itoa(id)
			if iStr[:digitLen/2] == iStr[digitLen/2:] {
				invalidIdsSum += id
				println("Invalid id : ", id)
			}

			id = roundUpToEvenLength(id + 1)
		}
	}
	println(invalidIdsSum)
}
