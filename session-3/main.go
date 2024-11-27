package main

import (
	"fmt"
)

func main() {
	var quantity int
	var lenght, width float64
	var customerName string

	lenght, width = 1.2, 2.4

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(lenght*width, "square meters")
}
