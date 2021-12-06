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
	nums := coding.ParseListAsInt(strings.Split(input[0], ","))
	for days := 0; days < 256; days++ {
		for idx, fish_age := range nums {
			if fish_age == 0 {
				nums[idx] = 6
				nums = append(nums, 8)
			} else {
				nums[idx]--
			}
		}
	}
	return len(nums)
}
