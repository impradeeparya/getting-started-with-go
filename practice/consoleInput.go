package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered text : ", input)

	fmt.Print("Enter number : ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println("error while parsing number : ", err)
	} else {
		fmt.Println("Entered number ", aFloat)
	}
}
