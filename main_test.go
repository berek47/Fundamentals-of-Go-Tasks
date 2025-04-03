package main

import (
<<<<<<< HEAD
	"task2/palindrome_check"
	"task2/word_frequency"
	"testing"
)

func TestMainFunction(t *testing.T) {
    // Test Palindrome Check
    palindrome := "A man, a plan, a canal, Panama"
    if !palindrome_check.IsPalindrome(palindrome) {
        t.Errorf("Expected true for palindrome, but got false")
    }

    // Test Word Frequency
    text := "Hello, world! Hello Go, hello!"
    expected := map[string]int{
        "hello": 3,
        "world": 1,
        "go":    1,
    }
    result := word_frequency.WordFrequency(text)
    for word, expectedCount := range expected {
        if result[word] != expectedCount {
            t.Errorf("For word %s, expected %d but got %d", word, expectedCount, result[word])
        }
    }
=======
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
>>>>>>> 6bcaf91a95b3d54fe1753f232ff03b04f4906ca7
}
