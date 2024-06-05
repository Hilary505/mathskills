package math_skills

import (
	"fmt"
	"os"
)

// function calculates the average
func Average(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("The data file is empty, enter some data")
		os.Exit(1)
	}
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	average := sum / float64(len(data))
	return average
}
