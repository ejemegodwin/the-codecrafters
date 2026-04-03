package main

import (
	"fmt"
)

func main() {
	for {
		var num1 int
		var num2 int
		fmt.Println("enterfirst number: ")
		fmt.Scanln(&num1)
		fmt.Println("enter second number: ")
		fmt.Scanln(&num2)
		var operations string
		fmt.Println("chose an opration (+ | - | * | / | help | quit)")
		fmt.Scanln(&operations)
		switch operations {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			fmt.Println(num1 / num2)
		case "help":
			fmt.Println("Choose from available Operations; add, sub, mul, div, help, quit")
		case "quit":
			fmt.Println("GoodBye!!!")
			return
		default:
			fmt.Println("Input a valid Expression")
		}
	}
}
