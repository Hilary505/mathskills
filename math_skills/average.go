package math_skills

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function calculates the average
func Average(str string) float64 {
	data := strings.Split(str, "\n")
	var sum float64

	for _, v := range data {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("An error:", err)
			os.Exit(1)
		}
		sum += num
	}
	var average = (sum) / float64(len(data))
	return average
}
