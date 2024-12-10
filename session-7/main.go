package main

func main() {
	aufgabe1()
	aufgabe2()
	aufgabe3()
}

func aufgabe1() {
	for i := 0; i <= 100000; i++ {
		if i%5 == 0 {
			println(i)
		}
	}
}

func aufgabe2() {
	alter := 14
	if alter < 16 {
		println("Du darfst keinen Alkohol trinken.")
	} else if alter == 16 || alter == 17 {
		println("Du darfst Bier und Wein trinken.")
	} else {
		println("Du darfst alles trinken.")
	}
}

func aufgabe3() {
	number := 5

	result := 0
	for i := 0; i <= number; i++ {
		result = result + i // result += i
	}
	println(result)
}
