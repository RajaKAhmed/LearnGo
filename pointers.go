package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Memory Address of anInt:", p)
	fmt.Println("Value stored in memory location p:", *p)
}
