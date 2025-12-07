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

	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	inputMaxCols := 0
	for i := range input {
		inputMaxCols = max(inputMaxCols, len(input[i]))
	}

	for i := range input {
		for len(input[i]) < inputMaxCols {
			input[i] += " "
		}
	}

	for i := range input {
		assert.Equals(len(input[i]), inputMaxCols)
	}

	operands := input[:len(input)-1]
	operandsTransposed := make([][]byte, inputMaxCols)
	for i := range operandsTransposed {
		operandsTransposed[i] = make([]byte, len(operands))
		for j := range operandsTransposed[i] {
			operandsTransposed[i][j] = input[j][i]
		}
		fmt.Printf("[%v]\n", string(operandsTransposed[i]))
	}

	operatorsStr := input[len(input)-1:][0]
	operators := strings.Fields(operatorsStr)

	grandTotal := 0

	exercises := [][]int{}
	exercise := []int{}
	for _, operand := range operandsTransposed {
		operandStr := strings.TrimSpace(string(operand))
		if (operandStr == "") {
			exercises = append(exercises, exercise)
			exercise = []int{}
		} else {
			operandNum, err := strconv.Atoi(operandStr)
			assert.NoError(err)
			exercise = append(exercise, operandNum)
		}
	}
	exercises = append(exercises, exercise)
	assert.Equals(len(exercises), len(operators))

	for i, operator := range operators {
		if operator == "+" {
			accumulator := 0
			for _, operand := range exercises[i] {
				accumulator += operand
			}
			grandTotal += accumulator
			fmt.Printf("Sum %d = %d\n", i, accumulator)
		} else {
			accumulator := 1
			for _, operand := range exercises[i] {
				accumulator *= operand
			}
			grandTotal += accumulator
			fmt.Printf("Mul %d = %d\n", i, accumulator)
		}
	}

	fmt.Printf("Grand total = %d\n", grandTotal)
}

