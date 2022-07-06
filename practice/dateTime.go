package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println("Current time ", now)

	var currentTime = time.Now()
	fmt.Println("Current time ", currentTime)

	dateTime := time.Date(1990, time.October, 19, 0, 0, 0, 0, time.Local)
	fmt.Println("date ", dateTime)

	parsedDateTime := dateTime.Format(time.ANSIC)
	fmt.Println("ANSIC date formatter ", parsedDateTime)
}
