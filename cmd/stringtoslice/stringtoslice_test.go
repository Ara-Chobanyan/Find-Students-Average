package stringtoslice

import (
	"testing"
)

func TestStringToSlice(t *testing.T) {
	strings := "55"

	got := ConvertStringToSlice(strings)
	want := []int{55}

	if got[0] != want[0] {
		t.Errorf("got %d want %d given, %v", got, want, strings)
	}
}
