package utils

import (
	"math"
)

func Abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func MeanLuminosity(arr []float64) float64 {
	mean := 0.0
	n := len(arr)
	for _, v := range arr {
		mean += v
	}
	return mean / float64(n)
}

func StdDev(arr []float64, mean float64) float64 {
	variance := 0.0
	n := len(arr)
	for _, v := range arr {
		variance += (v - mean) * (v - mean)
	}
	return math.Sqrt(variance / float64(n))
}

func CalcBounds(mean float64, stdDev float64) (lower, upper float64) {
	upper = mean + 2*stdDev
	lower = mean - 2*stdDev

	// fallback values if values do not lie in [0,1] interval
	if lower < 0 {
		mean = upper
		lower = 0
		upper += 2 * stdDev
	}
	if upper > 1 {
		upper = 1
	}
	return
}
