package main

import (
	"advent2021/utils/files"
	"bytes"
	"strconv"
)

func main() {
	input := files.ReadFile(3, "\n")
	println(solvePartX(input))
}

func solvePartX(input []string) int64 {
	bin_str_len := len(input[0])
	zeros, ones := make([]int, bin_str_len), make([]int, bin_str_len)
	for _, bin_str := range input {
		for bin_str_idx, char := range bin_str {
			switch char {
			case '0':
				zeros[bin_str_idx] += 1
			case '1':
				ones[bin_str_idx] += 1
			}
		}
	}
	var gamma, epsilon bytes.Buffer
	for i := 0; i < bin_str_len; i++ {
		if zeros[i] > ones[i] {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		} else {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		}
	}

	gamma_int, _ := strconv.ParseInt(gamma.String(), 2, 64)
	epsilon_int, _ := strconv.ParseInt(epsilon.String(), 2, 64)

	return gamma_int * epsilon_int

}
