package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	firstNumber := input()
	secondNumber := input()

	operator := operation()

	switch strings.TrimSpace(operator) {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println("Enter valid operations + - / *")
	}
}

func input() float64 {
	fmt.Print("Enter number ")
	value, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		panic(any("Enter number only"))
	} else {
		return number
	}
}

func operation() string {
	fmt.Print("Enter Operation ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		panic(any("Enter valid operations + - / *"))
	} else {
		return operator
	}
}
