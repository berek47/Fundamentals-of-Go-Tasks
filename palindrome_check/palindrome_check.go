package palindrome_check

import (
	"regexp"
	"strings"
)

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(input string) bool {
    input = strings.ToLower(input)
    re := regexp.MustCompile(`[^\w]`)
    input = re.ReplaceAllString(input, "")
    reversed := reverse(input)
    return input == reversed
}

// Helper function to reverse a string
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
