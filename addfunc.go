package main

import (
	"fmt"
)

func addme(x int, y int) int {
	return x + y
}

func memaddr(x *int) {
	fmt.Println(x)
	return
}

func main() {
	fmt.Println(addme(2, 4))
	t := addme(4, 5)
	s := "Hello World"
	fmt.Println(s, &s)
	fmt.Println(t, &t)
	i := 1
	fmt.Println(&i)
	memaddr(&i)
}
