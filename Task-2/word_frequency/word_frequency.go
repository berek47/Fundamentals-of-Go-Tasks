package word_frequency

import (
	"regexp"
	"strings"
)

// WordFrequency takes a string and returns a map with word frequencies
func WordFrequency(text string) map[string]int {
    text = strings.ToLower(text)
    re := regexp.MustCompile(`[^\w\s]`)
    text = re.ReplaceAllString(text, "")
    words := strings.Fields(text)

    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }

    return wordCount
}
