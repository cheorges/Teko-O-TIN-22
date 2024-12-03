package main

import (
	"fmt"
	"session-6/util"
)

/*
1. Zinsrechner: Entwickeln Sie einen Zinsrechner, der den Benutzer nach einem Anfangskapital,
Zinssatz und Laufzeit in Jahren fragt. Verwenden Sie eine For-Schleife, um das Kapital für
jedes Jahr zu berechnen und auszugeben. Der Benutzer soll entscheiden können, ob er eine
weitere Berechnung durchführen möchte.

Beispiel:
Willkommen zum Zinsrechner!
Geben Sie das Anfangskapital ein (in CHF): 10000
Geben Sie den jährlichen Zinssatz ein (%): 2.5
Geben Sie die Laufzeit in Jahren ein: 5

Berechnung 1:
Jahr 1: 10250.00 CHF
Jahr 2: 10506.25 CHF
Jahr 3: 10768.91 CHF
Jahr 4: 11038.13 CHF
Jahr 5: 11314.08 CHF

Möchten Sie eine weitere Berechnung durchführen? (j/n): j

Geben Sie das Anfangskapital ein (in CHF): 5000
Geben Sie den jährlichen Zinssatz ein (%): 1.8
Geben Sie die Laufzeit in Jahren ein: 3

Berechnung 2:
Jahr 1: 5090.00 CHF
Jahr 2: 5181.62 CHF
Jahr 3: 5274.89 CHF

Möchten Sie eine weitere Berechnung durchführen? (j/n): n
Vielen Dank für die Nutzung des Zinsrechners!
*/

func maain() {
	fmt.Println("Willkommen zum Zinsrechner!")

	var calculationCount int

	for {
		calculationCount++

		initialCapital := util.GetNumberFromUser("Geben Sie das Anfangskapital ein (in CHF):")
		interestRate := util.GetNumberFromUser("Geben Sie den jährlichen Zinssatz ein (%): ")
		years := util.GetNumberFromUserAsInt("Geben Sie die Laufzeit in Jahren ein: ")

		fmt.Printf("Berechnung: %d\n", calculationCount)

		calculateAndPrintInterest(initialCapital, interestRate, years)

		shouldContinue := util.GetStringFromUser("Möchten Sie eine weitere Berechnung durchführen? (j/n): ")

		if shouldContinue != "j" {
			break
		}

	}

	fmt.Println("Vielen Dank für die Nutzung des Zinsrechners!")

}

func calculateAndPrintInterest(initialCapital float64, interestRate float64, years int) {
	capital := initialCapital

	for i := 1; i <= years; i++ {
		capital = capital * (1 + interestRate/100)
		fmt.Printf("Jahr %d: %.2f CHF\n", i, capital)
	}
	fmt.Println()
}
