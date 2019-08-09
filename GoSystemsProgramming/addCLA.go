package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0
	for i := i; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum += temp
	}
	fmt.Println("Sum:", sum)
}
