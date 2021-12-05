package main

import (
	"advent2021/utils/coding"
	"advent2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(5, "\n")
	println(solvePartX(input))
}

type Coord struct {
	X, Y int
}

var coord_map = make(map[Coord]int)

func increment(coord Coord) {
	_, ok := coord_map[coord]
	if ok {
		coord_map[coord]++
	} else {
		coord_map[coord] = 1
	}
}

func parse_line(line string) (coord_list []Coord) {
	var half1, half2 string
	for idx, char := range line {
		if char == '-' {
			half1, half2 = line[:idx-1], line[idx+3:]
			break
		}
	}
	split_half_1, split_half_2 := coding.ParseListAsInt(strings.Split(half1, ",")), coding.ParseListAsInt(strings.Split(half2, ","))
	x1, y1, x2, y2 := split_half_1[0], split_half_1[1], split_half_2[0], split_half_2[1]
	switch {
	case x1 == x2:
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		coord_list = make([]Coord, 0, y2-y1+1)
		for start := y1; start <= y2; start++ {
			coord_list = append(coord_list, Coord{x1, start})
		}
	case y1 == y2:
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		coord_list = make([]Coord, 0, x2-x1+1)
		for start := x1; start <= x2; start++ {
			coord_list = append(coord_list, Coord{start, y1})
		}
	default:
		if x2 < x1 {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		dir := 1
		if y2 < y1 {
			dir = -1
		}
		coord_list = make([]Coord, 0, x2-x1+1)
		for delta := 0; delta <= x2-x1; delta++ {
			coord_list = append(coord_list, Coord{x1 + delta, y1 + delta*dir})
		}
	}
	return
}

func solvePartX(input []string) (count int) {
	lines := make([][]Coord, 0, len(input))
	for _, str := range input {
		coords := parse_line(str)
		lines = append(lines, coords)
	}
	for _, line := range lines {
		for _, coord := range line {
			increment(coord)
		}
	}
	for _, coord_count := range coord_map {
		if coord_count > 1 {
			count++
		}
	}
	return count
}
