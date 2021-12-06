package main

import (
	"advent2021/utils/coding"
	"advent2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(6, "\n")
	println(solvePartX(input))
}

func solvePartX(input []string) int {
	var fish_age_counts = make(map[int]int, 0)
	nums := coding.ParseListAsInt(strings.Split(input[0], ","))
	for _, fish_age := range nums {
		fish_age_counts[fish_age]++
	}
	for i := 0; i < 256; i++ {
		var zero_count_to_add int
		for day := 0; day <= 8; day++ {
			switch day {
			case 0:
				zero_count_to_add = fish_age_counts[day]
				fish_age_counts[0] = 0
			default:
				fish_age_counts[day-1] += fish_age_counts[day]
				fish_age_counts[day] = 0
				if day == 6 || day == 8 {
					fish_age_counts[day] = zero_count_to_add
				}
			}
		}
	}
	count := 0
	for _, val := range fish_age_counts {
		count += val
	}
	return count
}
