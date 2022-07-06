package main

import (
	"fmt"
)

const aConstant = "This is constant"
const aIntegerConstant = 10

//myString := "This is my string" // not allowed here
//const myString := "This is my string" // not allowed here

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("Variable type is %T\n", aString)

	var anInteger int = 69
	fmt.Println(anInteger)
	fmt.Printf("Variable type is %T\n", anInteger)

	var defaultInteger int
	fmt.Println(defaultInteger)

	myString := "This is my string"
	fmt.Println(myString)
	fmt.Printf("myString type is %T\n", myString)

	fmt.Println(aConstant)
	fmt.Printf("Constant type is %T\n", aConstant)

	fmt.Println(aIntegerConstant)
	fmt.Printf("Constant type is %T\n", aIntegerConstant)
}
