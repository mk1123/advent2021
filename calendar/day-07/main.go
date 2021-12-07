package main

import (
	"advent2021/utils/coding"
	"advent2021/utils/files"
	"math"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(7, "\n")
	println(solvePart2(input))
}

func abs_int(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func solvePart1(input []string) int {
	nums := coding.ParseListAsInt(strings.Split(input[0], ","))
	median := median(nums)
	diff := 0
	for _, num := range nums {
		diff += abs_int(num - median)
	}
	return diff
}

func median(list []int) int {
	sort.Ints(list)
	if len(list)%2 == 0 {
		return (list[len(list)/2] + list[(len(list)+1)/2]) / 2
	} else {
		return list[len(list)/2]
	}
}

func triangle_cost(nums []int, u int) (cost int) {
	for _, num := range nums {
		diff := abs_int(num - u)
		cost += diff * (diff + 1) / 2
	}
	return
}

func solvePart2(input []string) (min_diff_sum int) {
	nums := coding.ParseListAsInt(strings.Split(input[0], ","))
	min_diff_sum = math.MaxInt64
	for num := coding.MinList(nums); num <= coding.MaxList(nums); num++ {
		curr_diff := triangle_cost(nums, num)
		if curr_diff < min_diff_sum {
			min_diff_sum = curr_diff
		}
	}
	return
}
