package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func main() {
	rand2.Seed(time.Now().UnixMilli())

	day := rand2.Intn(7) + 1

	switch day {
	case 1:
		fmt.Println("It's Sunday")
	case 2:
		fmt.Println("It's Monday")
	case 3:
		fmt.Println("It's Tuesday")
	case 4:
		fmt.Println("It's Wednesday")
	case 5:
		fmt.Println("It's Thursday")
	case 6:
		fmt.Println("It's Friday")
	case 7:
		fmt.Println("It's Saturday")
		fallthrough
	default:
		fmt.Println("Some other day")
	}
}
