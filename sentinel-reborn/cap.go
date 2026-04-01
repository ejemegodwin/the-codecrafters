package main

import (
	"strings"
)

func ParseModifier(t string) string {
	s := strings.Fields(t)
	result := []string{}

	for i := 0; i < len(s); i++ {
		word := s[i]
		switch word {
		case "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}

		case "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}

		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = strings.Title(result[len(result)-1])
			}
		default:
			result = append(result, word)
		}

	}
	return strings.Join(result, " ")
}
