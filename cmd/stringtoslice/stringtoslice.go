package stringtoslice

import (
	"strconv"
	"strings"
)

// ConvertStringToSlice - Converts strings of numbers into []int
func ConvertStringToSlice(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}
