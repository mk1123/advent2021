package main

import (
	"advent2021/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	println(solvePartX(input))
}

func solvePartX(input []string) int {
	vert, horiz := 0, 0

	for _, direction_str := range input {
		split_direction_str := strings.Fields(direction_str)
		parsed_num, _ := strconv.Atoi(split_direction_str[1])
		switch split_direction_str[0] {
		case "forward":
			horiz += parsed_num
		case "down":
			vert += parsed_num
		case "up":
			vert -= parsed_num
		}
	}

	return vert * horiz
}
