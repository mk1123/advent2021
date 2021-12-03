package main

import (
	"advent2021/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(3, "\n")
	println(solvePartX(input))
}

func most_common(bin_strs []string, idx int, isOxy bool) bool {
	common_flag := 0
	for _, bin_str := range bin_strs {
		switch bin_str[idx] {
		case '0':
			common_flag--
		case '1':
			common_flag++
		}
	}
	if isOxy {
		return common_flag >= 0
	} else {
		return common_flag < 0
	}
}

var b2c = map[bool]byte{false: '0', true: '1'}

func solvePartX(input []string) int64 {
	oxy_arr := append(make([]string, 0, len(input)), input...)
	co2_arr := append(make([]string, 0, len(input)), input...)
	var oxy_bool, co2_bool bool

	for string_index := 0; len(oxy_arr) > 1; string_index++ {
		oxy_bool = most_common(oxy_arr, string_index, true)
		idx := 0
		for _, bin_str := range oxy_arr {
			if bin_str[string_index] == b2c[oxy_bool] {
				oxy_arr[idx] = bin_str
				idx++
			}
		}
		oxy_arr = oxy_arr[:idx]
	}

	oxy, _ := strconv.ParseInt(oxy_arr[0], 2, 64)

	for string_index := 0; len(co2_arr) > 1; string_index++ {
		co2_bool = most_common(co2_arr, string_index, false)
		idx := 0
		for _, bin_str := range co2_arr {
			if bin_str[string_index] == b2c[co2_bool] {
				co2_arr[idx] = bin_str
				idx++
			}
		}
		co2_arr = co2_arr[:idx]
	}

	co2, _ := strconv.ParseInt(co2_arr[0], 2, 64)

	return oxy * co2
}
