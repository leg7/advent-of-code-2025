package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"leg7.com/aoc2025/utils/assert"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	assert.True(scanner.Scan(), "No input provided")

	fields := strings.Fields(scanner.Text())
	fieldCount := len(fields)

	calculations := make([][]int, fieldCount)
	for i := range fieldCount {
		calculations[i] = make([]int, 0, 4)
	}

	for scanner.Scan() {
		for i := range fields {
			operand, err := strconv.Atoi(fields[i])
			assert.NoError(err)
			calculations[i] = append(calculations[i], operand)
		}

		fields = strings.Fields(scanner.Text())
		assert.Equals(len(fields), fieldCount)
	}
	fmt.Println(calculations)

	println(len(calculations[0]), cap(calculations[0]))

	grandTotal := 0
	operator := fields
	for i := range operator {
		operands := calculations[i]

		switch operator[i] {
		case "+": {
			accumulator := 0
			for j := range len(operands) {
				accumulator += operands[j]
			}
			fmt.Printf("Sum %d = %d\n", i, accumulator)
			grandTotal += accumulator
		}
		case "*": {
			accumulator := 1
			for j := range len(operands) {
				accumulator *= operands[j]
			}
			fmt.Printf("Mul %d = %d\n", i, accumulator)
			grandTotal += accumulator
		}
		}
	}

	fmt.Printf("\nGrand total = %d\n", grandTotal)
}

