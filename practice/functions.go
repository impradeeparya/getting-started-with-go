package main

import "fmt"

func main() {

	doSomething()
	sum := addValue(1, 2)
	fmt.Println(sum)

	sum, counts := addAll(1, 2, 3, 4, 5)
	fmt.Println(sum, counts)
}

func doSomething() {
	fmt.Println("doing something")
}

func addValue(value1 int, value2 int) int {
	return value2 + value1
}

func addAll(values ...int) (int, int) {
	var total int
	for _, value := range values {
		total += value
	}
	return total, len(values)
}
