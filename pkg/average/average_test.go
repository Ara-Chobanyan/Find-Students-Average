package average

import "testing"

func TestFindAverage(t *testing.T) {
	average := FindAverage([]int{33, 44, 55, 66})
	expected := float32(49.5)

	if average != expected {
		t.Errorf("expected %f but got %f", expected, average)
	}
}

func BenchmarkFindAverage(b *testing.B) {
	average := FindAverage([]int{33, 44, 55, 66})
	expected := float32(49.5)

	if average != expected {
		b.Errorf("expected %f but got %f", expected, average)
	}
}

//Seems like the type conversion does not hit performance at all, then again I am a noob so what do I know
