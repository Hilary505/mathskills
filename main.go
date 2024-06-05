package main

import (
	"fmt"
	"math"
	a "maths/math_skills"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("The program expects atleast 2 arguments and not more than 2")
		os.Exit(0)
	}

	myargs := os.Args[1]
	contents, err := os.ReadFile(myargs)
	if err != nil {
		fmt.Println("An error has just occurred", err)
		return
	}
	data := string(contents)
	var data1 []float64
	for _, val := range strings.Fields(data) {
		newval, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("Error converting data:", err)
			continue
		}
		data1 = append(data1, newval)
	}

	//call other functions  to main for execution
	findAverage := a.Average(data1)
	fmt.Println("Average:", math.Round(findAverage))
	findMedian := a.FindMedian(data1)
	fmt.Println("Median:", math.Round(findMedian))
	findVariance := a.Variance(data1, findAverage)
	fmt.Printf("Variance: %.0f\n", math.Round(findVariance))
	findStandardDeviation := a.StandardDeviation(findVariance)
	fmt.Println("Standard Deviation:", math.Round(findStandardDeviation))

}
