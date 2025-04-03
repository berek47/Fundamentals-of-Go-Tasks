package main

import (
	"task1/grades"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	gradesList := []float64{90, 80, 70, 100}
	expectedAverage := 85.0

	result := grades.CalculateAverage(gradesList)
	if result != expectedAverage {
		t.Errorf("Expected %.2f but got %.2f", expectedAverage, result)
	}
}
