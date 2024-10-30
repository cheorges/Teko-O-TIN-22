package main

import "fmt"

func main() { // This is the main function and the entry point of the program
	fmt.Println("Main function is called")
	sayHello() // Call the sayHello function
	defaultVariableValues()
	variables()
	fmt.Println("Main function is about to exit")
}

func variables() {
	variable1 := 10    // Declare variable1 and assign value 10 to it
	var variable2 = 10 // Declare variable2 and assign value 10 to it
	var variable3 int  // Declare variable3 of type int
	variable3 = 10     // Assign value 10 to variable3

	fmt.Println("variable1:", variable1)
	fmt.Println("variable2:", variable2)
	fmt.Println("variable3:", variable3)

	variable1 = 20 // Reassign value 20 to variable1

	fmt.Println("variable1:", variable1)

	// variable1 = "20" // Does not compile because illegal data type for variable1!
}

func defaultVariableValues() {
	var variable1 int     // Default value of int is 0
	var variable2 float64 // Default value of float64 is 0
	var variable3 string  // Default value of string is ""
	var variable4 bool    // Default value of bool is false

	fmt.Println("variable1:", variable1)
	fmt.Println("variable2:", variable2)
	fmt.Println("variable3:", variable3)
	fmt.Println("variable4:", variable4)
}

// sayHello is a function that prints "Hello, World!" to the console
func sayHello() {
	fmt.Println("Hello, World!")
}

// This function is never called, so the IDE will show a warning and use other highlighting
func thisIsNeverUsed() {
	fmt.Println("This function is never called")
}
