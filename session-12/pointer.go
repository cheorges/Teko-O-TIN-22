package main

import "fmt"

func anwendungsfall() {
	var meinUserInput string // string

	fmt.Scan(&meinUserInput)

	fmt.Println(meinUserInput)
}

func examplePointer() {
	var number int // Datentype Ganzzahl (int) | default Wert 0

	number = 100

	fmt.Println(number, &number) // Wert (100), Speicherort des Werts

	var ptrNumber *int // Datentype Pointer Ganzzahl (*int) | default Wert nil

	ptrNumber = &number // ptrNumber zeigt auf die gleiche Adresse wie number

	fmt.Println(ptrNumber)

	number = 50

	fmt.Println("number", number, "number addr", &number, "ptrNumber", *ptrNumber, "ptrNumber addr", ptrNumber)
}
