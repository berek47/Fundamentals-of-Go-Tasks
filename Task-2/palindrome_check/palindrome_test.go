package palindrome_check // Same package as palindrome_check.go

import "testing"

func TestIsPalindrome(t *testing.T) {
    palindrome := "A man, a plan, a canal, Panama"
    if !IsPalindrome(palindrome) {
        t.Errorf("Expected true for palindrome, but got false")
    }

    nonPalindrome := "Hello, World!"
    if IsPalindrome(nonPalindrome) {
        t.Errorf("Expected false for non-palindrome, but got true")
    }
}
