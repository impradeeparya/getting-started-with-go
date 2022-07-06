package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 1, 2, 4
	iSum := i1 + i2 + i3
	fmt.Println("Sum is ", iSum)

	f1, f2, f3 := 1.59, 69.69, 34.234
	fSum := f1 + f2 + f3
	fmt.Println("Float sum ", fSum)
	fmt.Println("Float sum round off ", math.Round(fSum))
}
