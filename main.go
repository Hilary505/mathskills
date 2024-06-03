package main

import (
	"fmt"
 a "maths/math_skills" // import package math_skills as an alias
	"os"
	"math"
)

func main() {
	// length of commandline arguments should be two
	if len(os.Args) != 2 {
		fmt.Println("The program expects atleast 2 arguments and not more than 2")
		os.Exit(0)
	}
	//access the first argument of commandline
	myargs := os.Args[1]
	contents, err := os.ReadFile(myargs)
	if err != nil {
		fmt.Println("An error has just occurred", err)
		return
	}
	data := string(contents)
	//call other functions  to main for execution
	findAverage := a.Average(data)
	fmt.Println("Average:", math.Round(findAverage))
	findMedian := a.FindMedian(data)
	fmt.Println("Median:", math.Round(findMedian))
	findVariance := a.variance(data)
	fmt.Println("Variance:",math.Round(findVariance))
}
