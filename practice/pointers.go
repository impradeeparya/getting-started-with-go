package main

import "fmt"

func main() {

	valueOne := 32
	var firstPointer = &valueOne
	fmt.Println("ValueOne pointer ", *firstPointer)
	fmt.Println("ValueOne address ", firstPointer)

	valueTwo := 42
	var secondPointer = &valueTwo
	fmt.Println("ValueTwo pointer ", *secondPointer)
	fmt.Println("ValueTwo address ", secondPointer)

	*firstPointer = *secondPointer / 2
	fmt.Println("ValueOne pointer ", *firstPointer)
	fmt.Println("ValueOne address ", firstPointer)
	fmt.Println("ValueTwo pointer ", *secondPointer)
	fmt.Println("ValueTwo address ", secondPointer)
}
