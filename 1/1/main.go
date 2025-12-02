package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

func BoolToInt64(b bool) int64 {
    return int64(*(*byte)(unsafe.Pointer(&b)))
}

func Mod(a, b int64) int64 {
	remainder := a % b
	isNegativeMask := remainder >> 63
	return remainder + (isNegativeMask & b)
}

func main() {
	var dial, dialMax, total int64 = 50, 100, 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()

		ammount, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if (err != nil) {
			panic(err)
		}

		oldDial := dial
		direction := string(line[0])
		sign := 1 - 2 * BoolToInt64(direction == "L")
		dial = Mod(dial + sign*ammount, dialMax)

		total += BoolToInt64(dial == 0)
		fmt.Printf("%5d + %s%-5d = %-5d, total = %d\n", oldDial, direction, ammount, dial, total)
	}

	fmt.Println("Total = ", total)
}

