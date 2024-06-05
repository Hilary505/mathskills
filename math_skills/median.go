package math_skills

import "sort"

func FindMedian(data []float64) float64 {
	float := make([]float64, len(data))
	for i, s := range data {
		float[i] += s
	}
	sort.Float64s(float)

	var median float64
	if len(float)%2 == 1 {
		median = float[(len(float))/2]
	} else {
		median = (float[(len(float)/2)-1] + float[(len(float))/2]) / 2
	}
	return median
}
