package main

import (
	"advent2021/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(1, "\n")
	println(solvePartX(input))
}

func solvePartX(input []string) int {
	result := 0
	int_input := make([]int, len(input))
	for i, val := range input {
		int_input[i], _ = strconv.Atoi(val)
	}
	prev := int_input[0] + int_input[1] + int_input[2]
	var curr_sum int
	for i := 3; i < len(input); i++ {
		curr_sum = prev - int_input[i-3] + int_input[i]
		if curr_sum > prev {
			result++
		}
		prev = curr_sum
	}
	return result
}
