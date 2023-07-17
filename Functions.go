package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from the main function in Go!!")
	myPrintFunction()
	sum := addValues(1, 2)
	fmt.Println("Sum from addValues Function is: ", sum)
	multiSum, multiCount := addAllValues(1, 2, 3, 4, 5)
	fmt.Println("Sum of all values and Length of Array from addAllValues Function:", multiSum, multiCount)
}

func myPrintFunction() {
	fmt.Println("Hello from myPrintFunction function in Go!!!")
}

func addValues(v1 int, v2 int) int { // you can also declare it (v1,v2 int)
	return v1 + v2
}
func addAllValues(v1 ...int) (int, int) { //another way of declaration
	total := 0
	for _, v := range v1 {
		total += v
	}
	return total, len(v1)
}
