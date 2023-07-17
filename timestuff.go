package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I recorded this at: ", n)

	t := time.Date(2022, time.January, 01, 14, 48, 0, 0, time.UTC)
	fmt.Println("Go launched at: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sat Jan  1 14:48:00 2022")
	fmt.Printf("Type of Parsed time is %T\n", parsedTime)
}
