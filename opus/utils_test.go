package opus

import (
	"math"
)

// utility functions for unit tests

func addSineFloat32(buf []float32, sampleRate int, freq float64) {
	factor := 2 * math.Pi * freq / float64(sampleRate)
	for i := range buf {
		buf[i] += float32(math.Sin(float64(i) * factor))
	}
}

func addSine(buf []int16, sampleRate int, freq float64) {
	factor := 2 * math.Pi * freq / float64(sampleRate)
	for i := range buf {
		buf[i] += int16(math.Sin(float64(i)*factor) * (math.MaxInt16 - 1))
	}
}

func interleave(a []int16, b []int16) []int16 {
	if len(a) != len(b) {
		panic("interleave: buffers must have equal length")
	}
	result := make([]int16, 2*len(a))
	for i := range a {
		result[2*i] = a[i]
		result[2*i+1] = b[i]
	}
	return result
}
