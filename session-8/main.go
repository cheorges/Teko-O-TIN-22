package main

import "fmt"

func main() {
	greeting("Joe", "es")

	res := check("bla")
	fmt.Println(res)

	t, zahl := chain("a", "b", 1, 2)
	fmt.Println(t, zahl)
}

func greeting(name string, lang string) {
	if lang == "en" {
		fmt.Println("Hello", name)
	} else if lang == "de" {
		fmt.Println("Hallo", name)
	} else if lang == "es" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println(name)
	}
}

func check(password string) bool {
	secret := "mein gehimes password"
	return password == secret
}

func chain(text1 string, text2 string, num1 int, num2 int) (string, int) {
	text := text1 + text2
	num := num1 + num2
	return text, num
}
