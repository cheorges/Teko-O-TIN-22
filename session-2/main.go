package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Main Start")
	// Call functions ...
	fmt.Println("Main Ende")
}

func example3() {
	if 12 > 10 && 2 == 4 { // true && false = false
		fmt.Println("A")
	}

	if 12 >= 10 || 2 == 4 { // true || false = true
		fmt.Println("B")
	}

	if !(12 > 10 && 2 == 4) { // true && false = false -> true
		fmt.Println("C")
	}
}

func example2() {
	myBankAccount := 1200
	neededMoney := 500

	if myBankAccount >= neededMoney {
		fmt.Println("You can get the money")
		myBankAccount -= neededMoney // myBankAccount = myBankAccount - neededMoney
		fmt.Println("Your new bank account balance is:", myBankAccount)
	} else {
		fmt.Println("You cannot get the money")
		fmt.Printf("You have only %d money but you need %d\n", myBankAccount, neededMoney)
	}
}

func example1() {
	age := 18

	if age >= 16 {
		// Yes, can buy beer
		fmt.Println("Age:", age, "can buy beer")
	} else {
		// No, cannot buy beer
		fmt.Println("Age:", age, "cannot buy beer")
	}
}

func showName() {
	name := "John"
	fmt.Println("Name:", name)
}

func showVariableType() {
	var variable1 int
	var variable2 float64
	var variable3 string
	var variable4 bool

	fmt.Println("variable1:", reflect.TypeOf(variable1))
	fmt.Println("variable2:", reflect.TypeOf(variable2))
	fmt.Println("variable3:", reflect.TypeOf(variable3))
	fmt.Println("variable4:", reflect.TypeOf(variable4))
}
