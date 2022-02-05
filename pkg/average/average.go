package average

func FindAverage(grades []int) float32 {
	length := len(grades)
	sum := 0

	for i := 0; i < length; i++ {
		sum += grades[i]
	}
	return (float32(sum)) / (float32(length))
}

//O(n)
