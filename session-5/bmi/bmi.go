package bmi

import (
	"fmt"
	"session-5/util"
)

func calculateBMI(weight float64, height float64) float64 {
	return weight / ((height / 100) * (height / 100))
}

func printBMIStatus(bmi float64) {
	switch {
	case bmi < 18.5:
		fmt.Println("You are underweight.")
	case bmi >= 18.5 && bmi < 24.9:
		fmt.Println("You are normal weight.")
	case bmi >= 25 && bmi < 29.9:
		fmt.Println("You are overweight.")
	default:
		fmt.Println("You are obese.")
	}
}

func CalculateAndPrintOutBMI() {
	weight := util.GetNumberFromUser("Enter your weight in kg:")
	height := util.GetNumberFromUser("Enter your height in cm:")

	bmi := calculateBMI(weight, height)

	fmt.Println("Your BMI is", bmi)

	printBMIStatus(bmi)

}
