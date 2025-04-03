package main

import (
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
}
