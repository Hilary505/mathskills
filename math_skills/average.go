package math_skills

import (
	"fmt"
	"strconv"
	"strings"
)
// function that converts string to integer
func ConvertStringToInt(str string) int {
	Int_Average, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("An error occured", err)
		return 0
	}
	return int(Int_Average)
}
// function calculates the average 
func Average(str string) int {
	data1 := strings.Fields(str)
	sum := 0
	for _, ch := range data1 {
		num := ConvertStringToInt(ch)
		sum += num
	}
	var average = (sum) / len(data1)
	return average
}
