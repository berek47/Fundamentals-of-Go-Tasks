package grades

import "testing"

func TestCalculateAverage(t *testing.T) {
	// Test case with valid grades
	gradesList := []float64{90, 80, 70, 100}
	expectedAverage := 85.0

	result := CalculateAverage(gradesList)
	if result != expectedAverage {
		t.Errorf("Expected %.2f but got %.2f", expectedAverage, result)
	}
	
	emptyList := []float64{}
	expectedEmptyAverage := 0.0

	resultEmpty := CalculateAverage(emptyList)
	if resultEmpty != expectedEmptyAverage {
		t.Errorf("Expected %.2f but got %.2f for empty list", expectedEmptyAverage, resultEmpty)
	}
}
