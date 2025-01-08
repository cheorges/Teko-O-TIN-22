package main

import "fmt"

// isPrime prüft, ob eine Zahl eine Primzahl ist
func isPrime(number int) bool {
	// Eine Zahl kleiner oder gleich 1 ist keine Primzahl
	if number <= 1 {
		return false
	}

	// Prüfe jeden möglichen Teiler von 2 bis (number-1)
	for teiler := 2; teiler < number; teiler++ {
		// Wenn die Zahl durch einen Teiler teilbar ist (Rest 0),
		// dann ist es keine Primzahl
		if number%teiler == 0 {
			return false
		}
	}

	// Wenn kein Teiler gefunden wurde, ist es eine Primzahl
	return true
}

// main ist die Hauptfunktion des Programms, die die Benutzereingabe verarbeitet
// und das Ergebnis der Primzahlprüfung ausgibt
func main() {
	// Fordere den Benutzer zur Eingabe einer Zahl auf
	fmt.Println("Nummer eingeben:")
	// Lies die Benutzereingabe ein
	input := readUserInput()
	// Zeige an, welche Zahl geprüft wird
	fmt.Println("Prüfe", input)
	// Prüfe ob die eingegebene Zahl eine Primzahl ist und gib das Ergebnis aus
	if isPrime(input) {
		fmt.Println(input, "ist eine Primzahl")
	} else {
		fmt.Println(input, "ist keine Primzahl")
	}
}

// readUserInput liest eine ganze Zahl von der Standardeingabe ein
// und gibt diese zurück
func readUserInput() int {
	// Deklariere Variable für die Benutzereingabe
	var input int
	// Lies eine Zahl von der Standardeingabe ein und speichere sie in input
	fmt.Scan(&input)
	// Gib die eingelesene Zahl zurück
	return input
}
