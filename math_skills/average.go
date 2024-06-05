package math_skills

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function calculates the average
func Average(str string) float64 {
     
	data := strings.Fields(str)
     if len(data) == 0.0 {
		fmt.Println("The data file is empty, enter some data")
		os.Exit(0)
	 }
	sum := 0.0

	for _, v := range data {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("An error:", err)
			os.Exit(1)
		}
		sum += num
	}
	average := sum / float64(len(data))
	return average
}
