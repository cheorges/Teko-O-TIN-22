package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetNumberFromUser(prompt string) float64 {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func GetNumberFromUserAsInt(prompt string) int {
	number := GetNumberFromUser(prompt)
	return int(number)
}

func GetStringFromUser(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}
