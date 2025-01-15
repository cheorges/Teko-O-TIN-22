package main

import "fmt"

func main() {
	number := 10 // 10
	add(number)
	fmt.Println("(1) Number in main", number) // 10

	number2 := 100
	fmt.Println("Pointer (*int)", &number2) // 0x14000102030
	addToPointer(&number2)
	fmt.Println("(2) Number in main", number2)

}

func addToPointer(i *int) {
	fmt.Println("Pointer (*int)", &i, *i) // 0x14000102030, 100
	*i = *i + 10
	fmt.Println("Number in addToPointer", *i)
}

func add(i int) { // i = 10
	i = i + 10                      // i = 20
	fmt.Println("Number in add", i) // 20
}
