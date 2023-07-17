package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Yellow"
	colors[2] = "Green"
	fmt.Println(colors)
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Number of colors", len(colors))
	fmt.Println("Number of numbers", len(numbers))
}
