package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	split := strings.Split(s, " ")
	acronym := ""
	
	for _, element := range split {
		acronym += getFirstChar(element)
	}

	return strings.ToUpper(acronym)
}

func getFirstChar(element string) string {
	for _, char := range element {
		if unicode.IsLetter(char) {
			return string([]rune{char})
		}
	}
	return ""
}
