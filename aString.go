package main

import (
	"fmt"
)

func main() {
	var aString string = "Hello from aString"
	fmt.Println(aString)
	fmt.Printf("aString is type %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)
	fmt.Printf("anInt is type %T\n", anInt)

	var anIntInfer = 45
	fmt.Println(anIntInfer)
	fmt.Printf("anIntInfer is type %T\n", anIntInfer)
}
