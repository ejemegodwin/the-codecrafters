package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("invalid file")
		return
	}

	txt := string(data)

	t1 := UpN(txt)
	t2 := CapN(t1)
	t3 := LowN(t2)
	t4 := Punctuations(t3)
	t5 := ParseModifier(t4)
	t6 := FixQuotes(t5)
	t7 := Vol(t6)
	t8 := BinToDec(t7)
	t9 := replaceHexWithDec(t8)

	result := t9

	err = os.WriteFile("output.txt", []byte(result), 0644)

	if err != nil {
		fmt.Println("empty file")
		return
	}
	fmt.Println(result)

}
