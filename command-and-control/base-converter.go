package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex(hext string) int64 {

	hexacon, err := strconv.ParseInt(hext, 16, 64)
	if err != nil {
		fmt.Println("invalid hex input.Try again...")
	}
	return hexacon
}

func bin(bina string) int64 {
	bindec, err := strconv.ParseInt(bina, 2, 64)
	if err != nil {
		fmt.Println("invalid binary input.Try again...")
	}
	return bindec
}

func dec(deci string) (string, string) {
	d, err := strconv.ParseInt(deci, 10, 64)
	if err != nil {
		fmt.Println("invalid decimal input.Try again...")
	}
	decima := strconv.FormatInt(d, 2)
	decim := strconv.FormatInt(d, 16)
	return decima, strings.ToUpper(decim)

}

func main() {

	var input string
	var conv1 string
	var conv string

	for {
		fmt.Println("BASE CONVERTER")
		fmt.Println("_____________________________________________")
		fmt.Println("converter Base List:")
		fmt.Println("1.base-hex 16, 2.base-bin 2, 3.base-dec-10")
		fmt.Println("Choose a converter Base")
		fmt.Scanln(&conv, &conv1, &input)
		switch conv1 {
		case "hex":
			fmt.Println("Result:", "=", hex(input))
			fmt.Println("Done")
		case "bin":
			fmt.Println("Result:", "=", bin(input))
			fmt.Println("Done")
		case "dec":
			bin1, hex2 := dec(input)
			fmt.Println("Binary:", bin1)
			fmt.Println("Hex:", hex2)
			fmt.Println("Done")
		case "quit":
			fmt.Println("Goodbye from Base Converter!!")
			return
		default:
			fmt.Println("invalid option: choose a valid convert base")
		}
	}
}
