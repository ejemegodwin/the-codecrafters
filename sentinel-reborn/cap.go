package main

import (
	"fmt"
	"strings"
)
 func toUpper(word string) string {
	return strings.ToUpper(word)
 }

 func toLower(word string) string {
	return strings.ToLower(word)
 }

 func toTitle(word string) string {
	return strings.Title(word)
 }

func main() {
	fmt.Println(toUpper("hello royalty"))
	fmt.Println(toLower("KING ODOPKOLOPKOLO"))
	fmt.Println(toTitle("onminyi andrew okala"))
}
