package main

import (
	"fmt"
	"math"
)

func main() {
	i, j, k := 10, 5, 7
	sumThem := i + j + k
	fmt.Println("SumThem is: ", sumThem)

	l, m, n := 10.35, 5.733, 7.95511
	sumThese := l + m + n
	fmt.Println("sumThese is: ", sumThese)

	sumThese = math.Round(sumThese*100) / 100
	fmt.Println("SumThese is now rounded to: ", sumThese)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
