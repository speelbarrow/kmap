package kmap

import (
	"strconv"
	"strings"
)

// Parse takes an input string and converts it to arguments that can then be used to create a Kmap.
func Parse(input, delim string) ([]int, error) {
	var r []int

	for _, v := range strings.Split(input, delim) {
		if i, e := strconv.Atoi(v); e != nil {
			return nil, e
		} else {
			r = append(r, i)
		}
	}

	return r, nil
}
