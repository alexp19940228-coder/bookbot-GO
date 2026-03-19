package stats

import (
	"fmt"
	"strings"
	"unicode"
)

func WordCount(text string) string {
	splitString := strings.Fields(text)
	count := 0

	for i := 0; i < len(splitString); i++ {
		count++
	}

	return fmt.Sprintf("Found %d total words", count)
}

func CountChar(text string) map[rune]int {
	count := make(map[rune]int)
	textLower := strings.ToLower(text)

	for _, char := range textLower {
		if unicode.IsLetter(char) {
			count[char]++
		}

	}
	return count
}
