package main

import "fmt"

func main() {

	husky := Dog{"husky", 2}
	fmt.Println(husky)
	fmt.Printf("%+v\n", husky)
	husky.breed = "siberian husky"
	fmt.Printf("%+v\n", husky)

}

// Dog is a struct
type Dog struct {
	breed string
	age   int
}
