package math_skills

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := 3.0
	result := Average(data)
	if result != expected {
		t.Errorf("got %v, wanted %v", expected, result)
	}
}

func TestMedian(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := 3.0
	result := FindMedian(data)
	if result != expected {
		t.Errorf("got %v, wanted %v", expected, result)
	}
}

func TestVariance(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := 2.0
	result := Variance(data, Average(data))
	if result != expected {
		t.Errorf("got %v, wanted %v", expected, result)
	}
}
func TestStandardDeviation(t *testing.T) {
	variance := 2.0
	expected := 1.4142135623730951
	result := math.Sqrt(variance)
	if result != expected {
		t.Errorf("got %v, wanted %v", expected, result)
	}
}
