package coding

import (
	"math"
	"sort"
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

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func SliceContainsInt(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func DigitsToNum(digits []int) (num int) {
	for _, digit := range digits {
		num = num*10 + digit
	}
	return
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
