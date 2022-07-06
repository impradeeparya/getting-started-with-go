package main

import "fmt"

func main() {

	colors := []string{"Red", "Green", "Blue"}
	for index := 0; index < len(colors); index++ {
		fmt.Print(colors[index], " ")
	}

	fmt.Println()
	for index := range colors {
		fmt.Print(colors[index], " ")
	}

	fmt.Println()
	for _, value := range colors {
		fmt.Print(value, " ")
	}

	value := 1
	for value < 100 {
		if value > 10 {
			goto theEnd
		}
		fmt.Println(value)
		value++
	}

theEnd:
	fmt.Println(value)
}
