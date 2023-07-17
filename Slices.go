package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Yellow", "Green"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 123456
	numbers[1] = 12345
	numbers[2] = 1234
	numbers[3] = 123
	numbers[4] = 12
	numbers = append(numbers, 1)
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)
	fmt.Println("Number of colors", len(colors))
	fmt.Println("Number of numbers", len(numbers))

}
