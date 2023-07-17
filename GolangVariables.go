package main

import (
	"fmt"
)
const myconstInt uint = 3
func main() {
    var aString = "Hello for Linkedin"
	fmt.Println(aString)
	fmt.Printf("Variable type %T\n", aString)

	var anInt = 42
	fmt.Println(anInt)
	fmt.Printf("Variable type %T\n", anInt)
	fmt.Println(myconstInt)
	fmt.Printf("Variable type %T\n", myconstInt)
}