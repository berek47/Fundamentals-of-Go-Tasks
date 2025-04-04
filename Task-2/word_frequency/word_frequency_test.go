package word_frequency

import "testing"

func TestWordFrequency(t *testing.T) {
    text := "Hello, world! Hello Go, hello!"
    expected := map[string]int{
        "hello": 3,
        "world": 1,
        "go":    1,
    }

    result := WordFrequency(text)

    for word, expectedCount := range expected {
        if result[word] != expectedCount {
            t.Errorf("For word %s, expected %d but got %d", word, expectedCount, result[word])
        }
    }
}
