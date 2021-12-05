package main

import (
	"advent2021/utils/coding"
	"advent2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(4, "\n")
	println(solvePartX(input))
}

type bingo_card struct {
	marked            map[int]bool
	coords            map[int][]int
	row_mark_count    []int
	column_mark_count []int
}

func (card *bingo_card) mark_if_card_contains_num(num int) (finished bool) {
	coords, ok := card.coords[num]
	if ok {
		card.marked[num] = true
		row, col := coords[0], coords[1]
		card.row_mark_count[row] += 1
		card.column_mark_count[col] += 1
		if card.row_mark_count[row] == 5 || card.column_mark_count[col] == 5 {
			finished = true
		}
	}
	return finished
}

func (card *bingo_card) compute_score(called_out_num int) int {
	sum := 0
	for k := range card.coords {
		_, ok := card.marked[k]
		if !ok {
			sum += k
		}
	}
	return sum * called_out_num
}

func parse_comma_separated_string(str string) []int {
	sep_ints_as_strs := strings.Split(str, ",")
	return coding.ParseListAsInt(sep_ints_as_strs)
}

func parse_bingo_board(board_arr []string) *bingo_card {
	card := &bingo_card{
		make(map[int]bool),
		make(map[int][]int),
		make([]int, 5),
		make([]int, 5),
	}
	for row, row_str := range board_arr {
		row_vals := coding.ParseListAsInt(strings.Fields(row_str))
		for col, val := range row_vals {
			card.coords[val] = []int{row, col}
		}
	}
	return card
}

func solvePartX(input []string) int {
	first_line := input[0]
	called_nums := parse_comma_separated_string(first_line)
	var bingo_boards []*bingo_card
	for i := 2; i < 600; i += 6 {
		bingo_boards = append(bingo_boards, parse_bingo_board(input[i:i+5]))
	}
	num_incomplete := len(bingo_boards)
	for _, num := range called_nums {
		idx := 0
		for board_idx, board := range bingo_boards {
			if board.mark_if_card_contains_num(num) {
				num_incomplete -= 1
				if num_incomplete == 0 {
					return bingo_boards[board_idx].compute_score(num)
				}
			} else {
				bingo_boards[idx] = board
				idx++
			}
		}
		bingo_boards = bingo_boards[:idx]
	}
	return 0
}
