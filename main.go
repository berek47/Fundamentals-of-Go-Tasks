package main

import (
	"fmt"
	"task2/palindrome_check"
	"task2/word_frequency"
)

func main() {
    // Example for Palindrome Check
    palindromeText := "A man, a plan, a canal, Panama"
    if palindrome_check.IsPalindrome(palindromeText) {
        fmt.Println("The string is a palindrome!")
    } else {
        fmt.Println("The string is NOT a palindrome.")
    }

    // Example for Word Frequency Count
    text := "Hello, world! Hello Go, hello!"
    wordCount := word_frequency.WordFrequency(text)
    fmt.Println("Word Frequency Count:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
