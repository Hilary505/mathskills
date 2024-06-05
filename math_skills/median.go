package math_skills

import (
	"sort"
	"strconv"
	"strings"
)

// function that calculates the median based on lengths being odd or even
func FindMedian(str string) float64 {

	data := strings.Fields(str)
	float := make([]float64, len(data))
	for i, s := range data {
		value, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		float[i] += value
	}
	sort.Float64s(float)

	var median float64 // store the median value
	if len(float)%2 == 1 {
		median = float[(len(float))/2]
	} else {
		median = (float[(len(float)/2)-1] + float[(len(float))/2]) / 2
	}
	return median
}
