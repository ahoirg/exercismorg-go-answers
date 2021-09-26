// Package isogram is the solution to
// https://exercism.io/my/solutions/9280859a94724c47ac5be5a164d9d674
package isogram

import "strings"

func IsIsogram(text string) bool {
	text = strings.ToLower(text)

	for _, char := range text {
		if char != ' ' && char != '-' && strings.Count(text, string([]rune{char})) > 1 {
			return false
		}
	}
	return true
}
