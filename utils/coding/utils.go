package coding

import (
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
