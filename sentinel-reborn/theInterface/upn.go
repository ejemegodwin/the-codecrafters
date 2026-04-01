package theInterface

import (
	"regexp"
	"strings"
)

func UpN(input string) string {
	reCap := regexp.MustCompile(`\b(\w+) \(cap\)`)
	return reCap.ReplaceAllStringFunc(input, func(match string) string {
		parts := strings.Split(match, " ")
		word := parts[0]
		if len(word) == 0 {
			return match
		}
		return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	})
}