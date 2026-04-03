// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Ameh Destiny
// Squad:  the interface

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("SENTINEL STRING TRANSFORMER - ONLINE")
	fmt.Println("──────────────────────────────────────")

	for {
		fmt.Print("Enter command then text: ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		cmd, text := parseInput(input)

		if cmd == "exit" {
			fmt.Println("Oh!!, you are done... Goodbye.")
			return
		}

		if text == "" {
			fmt.Printf("No text provided. Usage: %s <text>\n", cmd)
			continue
		}

		switch cmd {
		case "upper":
			fmt.Println("Result:", toUpper(text))

		case "lower":
			fmt.Println("Result:", toLower(text))

		case "cap":
			fmt.Println("Result:", toCap(text))

		case "title":
			fmt.Println("Result:", toTitle(text))

		case "snake":
			fmt.Println("Result:", toSnake(text))

		case "reverse":
			fmt.Println("Result:", reverseWords(text))

		default:
			fmt.Printf("Unknown command: \"%s\"\n", cmd)
			fmt.Println("  Valid commands: upper, lower, cap, title, snake, reverse, exit")
		}
	}
}

func parseInput(input string) (string, string) {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return "", ""
	}

	cmd := strings.ToLower(parts[0])

	if len(parts) == 1 {
		return cmd, ""
	}

	text := strings.TrimSpace(input[len(parts[0]):])

	return cmd, text
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i, w := range words {
		runes := []rune(w)
		for a, b := 0, len(runes)-1; a < b; a, b = a+1, b-1 {
			runes[a], runes[b] = runes[b], runes[a]
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func toCap(text string) string {
	words := strings.Fields(text)

	for i, w := range words {
		runes := []rune(strings.ToLower(w))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func toUpper(text string) string {
	return strings.ToUpper(text)
}

func toLower(text string) string {
	return strings.ToLower(text)
}

func toSnake(s string) string {
	var result []rune

	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		} else if unicode.IsSpace(r) {
			result = append(result, '_')
		}
	}

	clean := []rune{}
	prevUnderscore := false

	for _, r := range result {
		if r == '_' {
			if !prevUnderscore {
				clean = append(clean, r)
			}
			prevUnderscore = true
		} else {
			clean = append(clean, r)
			prevUnderscore = false
		}
	}

	return strings.Trim(string(clean), "_")
}

func toTitle(s string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
		"is": true, "it": true,
	}

	words := strings.Fields(s)

	for i, w := range words {
		lower := strings.ToLower(w)

		if i == 0 || !smallWords[lower] {
			runes := []rune(lower)
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}

	return strings.Join(words, " ")
}
