package main

import "fmt"

func main() {
	motor1 := Motor{Leistung: 200} // instanzieren
	fmt.Println(motor1)

	motor2 := Motor{200, 5.5, 5} // bei allen werten können direcht werte übergeben werden
	fmt.Println(motor2)

	motor3 := Motor{200, 5.5, 5}

	if motor2 == motor3 {
		fmt.Println("SAME")
	}

	motor3.Leistung = 400

	fmt.Println(motor3.Leistung)
	fmt.Println(motor2.Leistung)
}

type Motor struct {
	Leistung       int
	Verbrauch      float64
	AnzahlZilinder int
}
