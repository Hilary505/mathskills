package math_skills

import "math"

func Variance(data []float64, avg float64) float64 {
	if len(data) == 0 {
		return 0
	}
	mean := avg
	variance := 0.0
	for _, value := range data {
		variance += (value - mean) * (value - mean)
	}

	return variance / float64(len(data))
}

func StandardDeviation(variance float64) float64 {
	sd := math.Sqrt(variance)
	return sd
}
