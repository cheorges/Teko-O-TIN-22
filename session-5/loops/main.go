package main

import "fmt"

func main() {

	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}

	fmt.Println("----------")

	for x := 0; x < 5; x++ {
		fmt.Println("x is now", x)
	}

	for x := 1; x < 1; x++ {
		fmt.Println("x is now", x)
	}

	/*
		for {
			fmt.Println("This will run forever")
		}
	*/

	i := 0
	for i < 5 {
		fmt.Println("i is now", i)
		i++
	}

	for x := 0; x < 5; x++ {
		fmt.Println("x is now", x)
	}

}
