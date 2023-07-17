package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1 // Original Implementation
	fmt.Println("Day", dow) // Only needed if the line above is uncommented to see what day.
	var result string
	switch dow {
	// switch dow := rand.Intn(7) + 1; dow { // Implementation on the same line.
	case 1:
		result = "It's Sunday"
		// fallthrough
	case 2:
		result = "It's Monday"

	case 3:
		result = "It's Tuesday"

	case 4:
		result = "It's Wednesday"

	case 5:
		result = "It's Thursday"

	case 6:
		result = "It's Jumma"

	default:
		result = "It's Sabbath"
	}
	fmt.Println(result)
}
