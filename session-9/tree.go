package main

import (
	"fmt"
	"strings"
)

func main() {
	height := 6 // Höhe des Baums
	printTree(height)
}

func printTree(height int) {
	// Zeichne die Baumkrone
	for i := 0; i < height; i++ {
		// Berechne die Anzahl der Sterne und Leerzeichen
		spaces := strings.Repeat(" ", height-i-1)
		stars := strings.Repeat("*", 2*i+1)

		// Drucke die aktuelle Zeile
		fmt.Println(spaces + stars)
	}

	// Zeichne den Stamm
	fmt.Println(strings.Repeat(" ", height-1) + "|")
}
