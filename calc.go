package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter 1st Number: ")
	reader := bufio.NewReader(os.Stdin)
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	fmt.Print("Enter 2nd Number: ")
	numInput1, _ := reader.ReadString('\n')
	aFloat1, err := strconv.ParseFloat(strings.TrimSpace(numInput1), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of 1st number", aFloat)
		fmt.Println("Value of 2nd number", aFloat1)
		fmt.Println("When we Add: ", aFloat+aFloat1)
		fmt.Println("When we Subtract: ", aFloat-aFloat1)
		fmt.Println("When we Divide: ", aFloat/aFloat1)
		fmt.Println("When we Multiply: ", aFloat*aFloat1)
	}
}
