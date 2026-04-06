package main

import (
	"fmt"
	"strconv"
)

func convertBase(num string, fromBase int, toBase int) (string, error) {
	decimalValue, err := strconv.ParseInt(num, fromBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(decimalValue, toBase), nil
}

func main() {
	num := "1E"
	fromBase := 16
	toBase := 10

	convertedNum, err := convertBase(num, fromBase, toBase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Converted Number:", convertedNum)
}
