package stringtoslice

import (
	"testing"
)

func TestStringToSlice(t *testing.T) {
	testCases := []struct {
		input  string
		output []int
	}{
		{"55", []int{55}},
		{"100", []int{100}},
		{"77", []int{77}},
		{"88", []int{88}},
		{"11", []int{11}},
		{"45", []int{45}},
		{"89", []int{89}},
		{"0", []int{0}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ConvertStringToSlice(tc.input)
			if actual[0] != tc.output[0] {
				t.Errorf("got %d; want %d", tc.output, actual)
			}
		})
	}

}
