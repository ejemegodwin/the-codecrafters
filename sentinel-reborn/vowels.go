package main

import (
	"strings"
)

func Vol(s string) string {
	word := strings.Fields(s)
	for i := range word {
		if word[i] == "a" && strings.ContainsAny(string(word[i+1][0]), "aeiouhAEIOUH") {
			word[i] = "an"
		} else if word[i] == "A" && strings.ContainsAny(string(word[i+1][0]), "aeiouhAEIOUH") {
			word[i] = "An"
		} else if word[i] == "an" && !strings.ContainsAny(string(word[i+1][0]), "aeiouhAEIOUH") {
			word[i] = "a"

		} else if word[i] == "An" && !strings.ContainsAny(string(word[i+1][0]), "aeiouhAEIOUH") {
			word[i] = "A"

		}
	}
	return strings.Join(word, " ")
}

