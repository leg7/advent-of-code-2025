package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"unsafe"
)

func BoolToInt64(b bool) int64 {
    return int64(*(*byte)(unsafe.Pointer(&b)))
}

func Abs(a int64) int64 {
	mask := a >> 63
	return (a ^ mask) - mask
}

func AbsBranchy(a int64) int64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Mod(a, b int64) int64 {
	remainder := a % b
	isNegativeMask := remainder >> 63
	return remainder + (isNegativeMask & b)
}

func ModBranchy(a, b int64) int64 {
	rem := a % b
	if rem < 0 {
		rem += b
	}
	return rem
}

func Branchless(reader io.Reader) int64 {
	var position, positionMax, total int64 = 50, 100, 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()

		ticks, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if (err != nil) {
			panic(err)
		}

		direction := string(line[0])
		ticksSign := 1 - 2 * BoolToInt64(direction == "L")

		revolutions := ticks / positionMax
		remainingTicks := ticks % positionMax
		positionPlusRemainingTicks := position + ticksSign*remainingTicks
		remainingTicksWrapsPosition := BoolToInt64(positionPlusRemainingTicks <= 0 && position != 0 || positionPlusRemainingTicks >= positionMax)
		total += revolutions + remainingTicksWrapsPosition

		// oldPosition := position
		position = Mod(position + ticksSign*ticks, positionMax)

		// fmt.Printf("%5d + %s%-5d = %-5d, total = %d\n", oldPosition, direction, ticks, position, total)
	}

	return total
}

func Branchy(reader io.Reader) int64 {
	var position, positionMax, total int64 = 50, 100, 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()

		ticks, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if (err != nil) {
			panic(err)
		}

		direction := string(line[0])
		newRawPosition := position
		if direction == "L" {
			newRawPosition -= ticks
		} else {
			newRawPosition += ticks
		}

		var revolutions int64 = 0
		if newRawPosition < 0 && position != 0 {
			revolutions += 1
		}
		revolutions += AbsBranchy(newRawPosition / positionMax)
		total += revolutions

		// oldPosition := position
		position = ModBranchy(newRawPosition, positionMax)

		if revolutions == 0 && position == 0 {
			total += 1
		}

		// fmt.Printf("%5d + %s%-5d = %-5d, total = %d\n", oldPosition, direction, ticks, position, total)
	}

	return total
}

func main() {
	println(Branchless(os.Stdin))
}

