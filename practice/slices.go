package main

import (
	"fmt"
	"sort"
)

func main() {

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Colors ", colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)

}
