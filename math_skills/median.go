package math_skills

import (
	"sort"
	"strconv"
	"strings"
)
// function that calculates the median based on lengths being odd or even
func FindMedian(str string) int {
	data1 := strings.Fields(str)
	ints := make([]int, len(data1))
	for i, s := range data1 {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] += num
	}
	// package sorts arranges data in descending or ascending order
	sort.Ints(ints)
	// var median is used to store the result of median
	var median int
	n := len(ints)
	if n%2 == 1 {
		median = ints[n/2]
	} else {
		median = (ints[n/2-1] + ints[n/2]) / 2
	}
	return median
}
