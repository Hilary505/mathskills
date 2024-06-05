package math_skills

import (
	"math"
)

// function that calculates variance  give slice of data
func Variance(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	// initialize sum to zero

	sum := 0.0
	for _, value := range data {
		sum += value
	}

	mean := sum / float64(len(data))
	variance := 0.0
	for _, value := range data {
		variance += (value - mean) * (value - mean)
	}

	return variance / float64(len(data))
}

// function that calculates standarddeviation
func StandardDeviation(variance float64) float64 {
	sd := math.Sqrt(variance)
	return sd
}
