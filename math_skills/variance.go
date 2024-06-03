package math_skills 

import (

)

func sumOfSquaredDifference(a  float64) float64{
	var sum float64
	for i , char := range data {
		sum = ((char[i] - a.Average)(char[i] - a.Average))
	}
	return sum
}

func Variance(sum   float64) float64{
	var variance float64
	variance = sum / len(data)
	return variance
}