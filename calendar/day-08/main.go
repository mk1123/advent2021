package main

import (
	"advent2021/utils/coding"
	"advent2021/utils/files"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	input := files.ReadFile(8, "\n")
	println(solvePart2(input))
}

func solvePart1(input []string) (count int) {
	unique_nums := []int{2, 3, 4, 7}
	for _, line := range input {
		input_output_split := strings.Split(line, " | ")
		line_outputs := strings.Split(input_output_split[1], " ")
		for _, word := range line_outputs {
			if coding.SliceContainsInt(unique_nums, len(word)) {
				count++
			}
		}
	}
	return
}

func ReverseMap(m map[int]string) map[string]int {
	n := make(map[string]int, len(m))
	for k, v := range m {
		n[coding.SortString(v)] = k
	}
	return n
}

func StringToSet(word string) mapset.Set {
	curr_word_set := mapset.NewSet()
	for _, char_str := range strings.Split(word, "") {
		curr_word_set.Add(char_str)
	}
	return curr_word_set
}

func solvePart2(input []string) (count int) {
	var num_to_string map[int]string
	var sorted_string_to_int map[string]int

	for _, line := range input {
		num_to_string = make(map[int]string)
		input_output_split := strings.Split(line, " | ")
		line_inputs := strings.Split(input_output_split[0], " ")
		line_outputs := strings.Split(input_output_split[1], " ")
		line_input_counts := make(map[int][]string)
		word_sets := make(map[string]mapset.Set)
		for _, word := range line_inputs {
			line_input_counts[len(word)] = append(
				line_input_counts[len(word)],
				word,
			)
			word_sets[word] = StringToSet(word)
		}
		// very crude CSP solver, lol
		num_to_string[1] = line_input_counts[2][0]
		num_to_string[7] = line_input_counts[3][0]
		num_to_string[4] = line_input_counts[4][0]
		num_to_string[8] = line_input_counts[7][0]

		// finding 3, 5, 2
		for _, word := range line_input_counts[5] {
			word_set := word_sets[word]
			switch {
			case word_set.Intersect(word_sets[num_to_string[7]]).Cardinality() == 3:
				num_to_string[3] = word
			case word_set.Intersect(word_sets[num_to_string[4]]).Cardinality() == 3:
				num_to_string[5] = word
			default:
				num_to_string[2] = word
			}
		}
		// finding 0, 6, 9
		for _, word := range line_input_counts[6] {
			word_set := word_sets[word]
			switch {
			case word_set.Intersect(word_sets[num_to_string[4]]).Cardinality() == 4:
				num_to_string[9] = word
			case word_set.Intersect(word_sets[num_to_string[1]]).Cardinality() == 1:
				num_to_string[6] = word
			default:
				num_to_string[0] = word
			}
		}

		// reverse map
		sorted_string_to_int = ReverseMap(num_to_string)
		output_digits := make([]int, 4)

		for i, word := range line_outputs {
			output_digits[i] = sorted_string_to_int[coding.SortString(word)]
		}
		count += coding.DigitsToNum(output_digits)
	}
	return
}
