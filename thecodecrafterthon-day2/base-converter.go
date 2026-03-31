package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter word to Convert: ")
		s.Scan()
		ie := strings.TrimSpace(s.Text())
		if ie == "quit" {
			return
		}
		run(ie)
	}
}

func run(ie string) {
	operation := strings.Fields(ie)
	if len(operation) != 2 {
		fmt.Println("Invalid Number")
		return
	}

	switch operation[1] {
	case "hex":
		toDec(operation[0], 16)

	case "bin":
		toDec(operation[0], 2)

	case "dec":
		fromDec(operation[0])

	case "help":
		fmt.Println("Valid options are hex | bin | dec | help | quit")

	default:
		fmt.Println("Invalid Operation")
	}
}

func toDec(a string, b int) {
	dec, err := strconv.ParseInt(a, b, 64)
	if err != nil {
		fmt.Println("Invalid Input my Guy")
		return
	}
	fmt.Println("Decimal:", dec)
}

func fromDec(num string) {
	del, e := strconv.ParseInt(num, 10, 64)
	if e != nil {
		fmt.Println("Invalid Input Champ")
		return
	}
	fmt.Println("Binary:", strconv.FormatInt(del, 2))
	fmt.Println("Hex:", strings.ToUpper(strconv.FormatInt(del, 16)))
}
