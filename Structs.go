// Structs or Data Structures. Java calls them classes. C and Golang calls them Structs. Structs are independent.
package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10} // Sort of like constructor in other languages but
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)                                        // Prints using printf the entire content of the values in struct
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight) //Use Printf  %v is known as verb. Prints on new line \n
	poodle.Weight = 9                                                  // Change values of Struct
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// Define your structure
type Dog struct {
	Breed  string
	Weight int
}
