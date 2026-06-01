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

func MeanLuminosity(arr [][]float64) float64 {
	mean := 0.0
	n := len(arr) * len(arr[0])
	for _, row := range arr {
		for _, val := range row {
			mean += val / float64(n)
		}
	}
	return mean
}

func StdDev(arr [][]float64, mean float64) float64 {
	variance := 0.0
	n := len(arr) * len(arr[0])
	for _, row := range arr {
		for _, val := range row {
			variance += math.Pow(val-mean, 2) / float64(n)
		}
	}
	return math.Sqrt(variance)
}
