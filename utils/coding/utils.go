package coding

import (
	"math"
	"strconv"
)

func ParseListAsInt(list []string) []int {
	ints := []int{}
	for _, i := range list {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ints = append(ints, j)
	}
	return ints
}

func SumList(list []int) (sum int) {
	for _, num := range list {
		sum += num
	}
	return sum
}

func MinList(nums []int) (min int) {
	min = math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return
}

func MaxList(nums []int) (max int) {
	max = math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}
