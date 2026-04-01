// package main

// import (
// 	"strconv"
// 	"strings"
// )

// func LowN(text string) string {
// 	opening_Marker := strings.Index(text, "(")
// 	closing_Marker := strings.Index(text, ")")
// 	marker := text[opening_Marker : closing_Marker+1]

// 	beforeMarker := text[:opening_Marker]
// 	words := strings.Fields(beforeMarker)

// 	num := 0
// 	for i := 0; i < len(marker)-1; i++ {
// 		if strings.ContainsAny(string(marker[i]), "123456789") {
// 			num, _ = strconv.Atoi(string(marker[i]))
// 		}
// 	}

// 	for i := len(words) - num; i < len(words); i++ {
// 		words[i] = strings.ToLower(words[i])
// 	}

// 	words = append(words, text[closing_Marker+1:])

// 	return strings.Join(words, " ")
// }

package main

import (
	"strconv"
	"strings"
)

func LowN(text string) string {
	newtext := []string{}
	txt := strings.Fields(text)
	for i := 0; i < len(txt); i++ {
		if i+1 < len(txt) && strings.HasPrefix(txt[i], "(low,") && strings.ContainsAny(txt[i], ",") {
			p := strings.TrimSuffix(txt[i+1], ")")
			m, err := strconv.Atoi(p)

			if err == nil {
				if m > len(newtext) {
					m = len(newtext)
				}
				for j := 1; j <= m; j++ {
					target := len(newtext) - j

					newtext[target] = strings.ToLower(newtext[target])
				}
				i++
				continue
			}

		}
		newtext = append(newtext, txt[i])

	}
	return strings.Join(newtext, " ")

}
