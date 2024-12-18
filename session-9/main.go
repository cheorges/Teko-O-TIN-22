package main

import "fmt"

func main() {
	// reference()
	// createSliceFromArray()
	// checkIteration()

	numbers := []int{5, 10, 15, 20, 25, 30} // Eindymanisches Array
	fmt.Println(numbers)

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix[1][1])
}

func checkIteration() {
	numbers := []int{5, 10, 15, 20, 25, 30}

	var numbersBiggerThan20 []int
	for i, number := range numbers {
		if number > 20 {
			fmt.Println(i, number)
			numbersBiggerThan20 = append(numbersBiggerThan20, number)
		}
	}
	fmt.Println(numbersBiggerThan20)
}

func createSliceFromArray() {
	letters := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(letters)

	lettersSlice := letters[1:3]
	fmt.Println(lettersSlice)
}

func reference() {
	primnumber := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primnumber)
	// notAllowed := [5]int{2, 3, 5, "bla", 11}

	weekdays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} // Array
	fmt.Println(weekdays)
	fmt.Println(weekdays[2]) // Index
	fmt.Println(weekdays[3]) // Index
	fmt.Println(weekdays[4]) // Index
	fmt.Println(weekdays[6]) // Index
	// fmt.Println(weekdays[7]) // Index out of range
	fmt.Println(len(weekdays))

	meineZahl := 42
	fmt.Println(meineZahl)
	meineZahl = 73
	// meineZahl = 3.14
	fmt.Println(meineZahl)

	weekdays[3] = "Little Friday"
	weekdays = [7]string{"Monday", "Tuesday", "Wednesday", "Little Friday", "Friday", "Saturday", "Sunday"}
	// weekdays = "Little Friday"
	fmt.Println(weekdays)

	// weekdays = append(weekdays, "New Day") // Geht nur bei Slices

	fmt.Println("-----------------------------")
	fmt.Println("Slices")
	fmt.Println("-----------------------------")

	weekdaysAsSlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} // Slice
	fmt.Println(weekdaysAsSlice)
	fmt.Println(weekdaysAsSlice[2]) // Index
	fmt.Println(len(weekdaysAsSlice))
	weekdaysAsSlice[3] = "Little Friday"
	fmt.Println(weekdaysAsSlice)

	weekdaysAsSlice = append(weekdaysAsSlice, "New Day")
	fmt.Println(weekdaysAsSlice, len(weekdaysAsSlice))
	fmt.Println(cap(weekdaysAsSlice))

	for i := 0; i < len(weekdaysAsSlice); i++ {
		fmt.Println(weekdaysAsSlice[i])
	}

	for i, day := range weekdaysAsSlice {
		fmt.Println(i, day)
	}
}
