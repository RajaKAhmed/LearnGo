package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Yellow"}
	fmt.Println(colors)
	// The following shows a three part loop. We initialize i to 0 processding with a semicolon; then we set i to less than the lenghth of colors
	// and increment i and print. This prints the slice values on a new line
	fmt.Println("For loop using a three part loop i=0; i less than the length of color; increment i++")
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	// The following usesi as well because we are using i in its own block of code, but here we use the range function to iterate through the slice it is much cleaner than i ++
	fmt.Println("For loop using Range Fucntion to iterate through a slice named colors")
	for i := range colors {
		fmt.Println(colors[i])
	}
	fmt.Println("A different implementation of For loop (the For each loop) with Range Fucntion to iterate through a slice named colors with omitting index the _")
	for _, color := range colors {
		fmt.Println(color)
	}
	// While Implementation in Go

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd //Go's implementation of break, continue....break and continue work as well this is an implementation of  goto function.
		}
	}
theEnd:
	fmt.Println("End of Program")
}
