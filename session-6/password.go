package main

import (
	"fmt"
	"session-6/util"
	"unicode"
)

/*
Passwort-Stärke-Prüfer:
Entwickeln Sie ein Programm, das den Benutzer auffordert,
ein Passwort einzugeben. Verwenden Sie eine For-Schleife, um jedes Zeichen des Passworts
zu überprüfen und die Stärke zu bewerten (z.B. Länge, Gross- und Kleinbuchstaben, Zahlen,
Sonderzeichen). Geben Sie am Ende eine Bewertung der Passwortstärke aus und lassen Sie
den Benutzer entscheiden, ob er ein weiteres Passwort prüfen möchte.

Beispiel:
Willkommen zum Passwort-Stärke-Prüfer!
Geben Sie ein Passwort ein: Passwort123

Passwort: Passwort123
Stärke: Mittel
- Enthält Gros- und Kleinbuchstaben
- Enthält Zahlen
- Länge ist ausreichend (mindestens 8 Zeichen)
Verbesserungsvorschlag: Fügen Sie Sonderzeichen hinzu

Möchten Sie ein weiteres Passwort prüfen? (j/n): j

Geben Sie ein Passwort ein: St@rk3sPassw0rt!

Passwort: St@rk3sPassw0rt!
Stärke: Sehr stark
- Enthält Gros- und Kleinbuchstaben
- Enthält Zahlen
- Enthält Sonderzeichen
- Länge ist sehr gut (mehr als 12 Zeichen)

Möchten Sie ein weiteres Passwort prüfen? (j/n): n

Auf Wiedersehen!
*/
func main() {
	fmt.Println("Willkommen zum Passwort-Stärke-Prüfer!")

	for {
		password := util.GetStringFromUser("Geben Sie ein Passwort ein: ")
		strength, feedback := checkPasswordStrength(password)

		fmt.Printf("\nPasswort: %s\n", password)
		fmt.Printf("Stärke: %s\n", strength)
		fmt.Println(feedback)
		fmt.Println()

		shouldContinue := util.GetStringFromUser("Möchten Sie eine weitere Berechnung durchführen? (j/n): ")

		if shouldContinue != "j" {
			break
		}
	}

}

func checkPasswordStrength(password string) (string, string) {
	lenght := len(password)
	hasUpper, hasLower, hasDigit, hasSpecial := false, false, false, false
	feedback := ""

	for i := 0; i < lenght; i++ {
		char := rune(password[i])
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}

	strenght := "Schwach"
	score := 0

	if hasUpper && hasLower {
		score++
		feedback += "- Enthält Gross- und Kleinbuchstaben\n"
	}

	if hasDigit {
		score++
		feedback += "- Enthält Zahlen\n"
	}

	if hasSpecial {
		score++
		feedback += "- Enthält Sonderzeichen\n"
	}

	if lenght >= 8 {
		score++
		if lenght >= 12 {
			score++
			feedback += "- Länge ist sehr gut (mehr als 12 Zeichen)\n"
		} else {
			feedback += "- Länge ist ausreichend (mindestens 8 Zeichen)\n"
		}
	}

	if score >= 5 {
		strenght = "Sehr stark"
	} else if score == 4 {
		strenght = "Stark"
	} else if score == 3 {
		strenght = "Mittel"
	}

	if !hasUpper || !hasLower {
		feedback += "Verbesserungsvorschlag: Fügen Sie Gross- und Kleinbuchstaben hinzu\n"
	}
	if !hasDigit {
		feedback += "Verbesserungsvorschlag: Fügen Sie Zahlen hinzu\n"
	}
	if !hasSpecial {
		feedback += "Verbesserungsvorschlag: Fügen Sie Sonderzeichen hinzu\n"
	}
	if lenght < 8 {
		feedback += "Verbesserungsvorschlag: Verwenden Sie mindestens acht Zeichen\n"
	}

	return strenght, feedback
}
