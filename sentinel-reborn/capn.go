package main

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalise(s string) string {
	a := strings.ToLower(s)
	runes := []rune(a)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func CapN(text string) string {
	newtext := []string{}
	txt := strings.Fields(text)
	for i := 0; i < len(txt); i++ {
		if i+1 < len(txt) && strings.HasPrefix(txt[i], "(cap,") && strings.ContainsAny(txt[i], ",") {
			p := strings.TrimSuffix(txt[i+1], ")")
			m, err := strconv.Atoi(p)

			if err == nil {
				if m > len(newtext) {
					m = len(newtext)
				}
				for j := 1; j <= m; j++ {
					target := len(newtext) - j

					newtext[target] = capitalise(newtext[target])
				}
				i++
				continue
			}

		}
		newtext = append(newtext, txt[i])

	}
	return strings.Join(newtext, " ")

}
