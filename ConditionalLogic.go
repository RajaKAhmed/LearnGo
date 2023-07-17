package main

import (
	"fmt"
)

func main() {
	someValue := 3
	var result string
	if someValue < 0 {
		result = "Less than Zero"
		fmt.Println(result, someValue)
	} else if someValue == 0 {
		result = "Equal to Zero"
		fmt.Println(result, someValue)
	} else if someValue > 0 {
		result = "Greater than Zero"
		fmt.Println(result, someValue)
	}

	if x := -3; x < 0 {
		result = "Less than Zero"
		fmt.Println(result, x)
	} else if x == 0 {
		result = "Equal to Zero"
		fmt.Println(result, x)
	} else {
		result = "Greater than Zero"
		fmt.Println(result, x)
	}

}
