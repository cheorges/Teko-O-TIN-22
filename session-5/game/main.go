package main

import (
	"fmt"
	"math/rand"
	"session-5/util"
)

/*
func main() {
	bmi.CalculateAndPrintOutBMI()
}
*/

func randomNumberFrom1To100() int {
	return rand.Intn(100) + 1
}

func main() {
	target := randomNumberFrom1To100()
	guess := util.GetNumberFromUserAsInt("Enter a number")

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else {
		fmt.Println("Oops. Your guess was HIGH.")
	}
}
