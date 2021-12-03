package main

import (
	"advent2021/utils/files"
	"log"
	"math"
	"strconv"
)

func main() {
	input := files.ReadFile(1, "\n")
	println(solvePart1(input))
}

func solvePart1(input []string) int {
	prev := math.MaxInt64
	result := 0
	for _, str := range input {
		str_val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		if str_val > prev {
			result++
		}
		prev = str_val
	}
	return result
}
