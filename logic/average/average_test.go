package average

import "testing"

func TestFindAverage(t *testing.T) {
	testCases := []struct {
		input []int
		want  float32
	}{
		{[]int{33, 44, 55, 66}, 49.5},
		{[]int{50, 100, 87, 66}, 75.75},
		{[]int{100, 100, 100, 100}, 100.0},
		{[]int{98, 77, 33, 78}, 71.5},
		{[]int{89, 51, 45, 100}, 71.25},
	}

	for _, tc := range testCases {
		t.Run("check average", func(t *testing.T) {
			actual := FindAverage(tc.input)
			if actual != tc.want {
				t.Errorf("got %f; want %f", actual, tc.want)
			}
		})
	}

}

func BenchmarkFindAverage(b *testing.B) {
	average := FindAverage([]int{33, 44, 55, 66})
	expected := float32(49.5)

	if average != expected {
		b.Errorf("expected %f but got %f", expected, average)
	}
}

//Seems like the type conversion does not hit performance in this scenario, then again I am a noob so what do I know
// To be honest I have a feeling that this bench mark test is proabay not even needed at all.
