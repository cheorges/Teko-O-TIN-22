# Selektion in Go Lang

## Theory

[If/Else and Conditional Statements in Go](https://dev.to/willvelida/ifelse-and-conditional-statements-in-go-2k9#:~:text=In%20Go%2C%20we%20can%20use,condition%20is%20true%20or%20false.)

### Mache dein Programm dynamischer 🚀

Benutzereingaben über die Konsole handlen in go:

```go
var value int
fmt.Print("Geben Sie eine Ganzahl ein: ")
fmt.Scan(&value)
```

ren:

1. `var value int`: Hier wird eine Variable namens value deklariert. Diese Variable ist vom Typ "int", was bedeutet, dass sie nur ganze Zahlen speichern kann. In diesem Fall wird die Variable verwendet, um die vom Benutzer eingegebene Zahl zu speichern.
2. `fmt.Print("Geben Sie eine Ganzahl ein: ")`: Diese Zeile gibt eine Aufforderung auf dem Bildschirm aus, die den Benutzer darum bittet, eine ganze Zahl einzugeben. Die Funktion Print aus dem Paket fmt wird verwendet, um diesen Text auf der Konsole anzuzeigen, ohne eine neue Zeile zu beginnen.
3. `fmt.Scan(&value)`: Hier erfasst der Code die Benutzereingabe. Die Funktion Scan aus dem Paket fmt wartet darauf, dass der Benutzer eine Zahl eingibt, und der eingegebene Wert wird dann der Variable value zugewiesen. Das &-Symbol vor value wird verwendet, um die Speicheradresse von value zu übergeben, damit die Funktion den eingegebenen Wert direkt an dieser Adresse ablegen kann.

In einfachen Worten: Der Codeblock fragt den Benutzer nach einer ganzen Zahl, nimmt die eingegebene Zahl entgegen und speichert sie in der Variable value. Diese Variable kann dann im Programm für weitere Berechnungen oder Entscheidungen verwendet werden.

Der grundlegende Ansatz kann auch für andere Variablentypen verwendet werden. Der Codeblock verwendet spezifisch den Datentyp int (für ganze Zahlen), aber du kannst ihn anpassen, um andere Datentypen zu akzeptieren. Zum Beispiel könntest du float64 für Gleitkommazahlen oder string für Zeichenketten verwenden.

Hier ist eine modifizierte Version des Codeblocks, die eine Gleitkommazahl einliest:

```go
package main

import "fmt"

func main() {
    var floatValue float64
    fmt.Print("Geben Sie eine Gleitkommazahl ein: ")
    fmt.Scan(&floatValue)

    fmt.Println("Eingegebene Zahl:", floatValue)
}
```

Und hier ist eine Version, die eine Zeichenkette einliest:

```go
package main

import "fmt"

func main() {
    var stringValue string
    fmt.Print("Geben Sie eine Zeichenkette ein: ")
    fmt.Scan(&stringValue)

    fmt.Println("Eingegebene Zeichenkette:", stringValue)
}
```

## Aufgaben

### 1. Gerade oder Ungerade

Schreiben Sie ein Programm, das eine ganze Zahl als Eingabe nimmt und ausgibt, ob sie gerade oder ungerade ist.

### 2. Notenrechner

Erstellen Sie ein Programm, das die Note eines Schülers als Eingabe entgegennimmt und die Note nach folgenden Kriterien ausgibt:

- 90-100: A
- 80-89: B
- 70-79: C
- 60-69: D
- 0-59: F

### 3. Altersklassifizierer

Schreiben Sie ein Programm, das das Alter einer Person als Eingabe nimmt und ausgibt, ob sie ein Kind, Jugendlicher, Erwachsener oder Senior ist.

### 4. Temperaturumrechner

Erstellen Sie ein Programm, das Temperaturen zwischen Fahrenheit und Celsius umrechnet. Fragen Sie den Benutzer nach der Temperatur und der Einheit (F oder C) und geben Sie dann die umgerechnete Temperatur aus.

### 5. Benutzername und Passwortprüfer

Schreiben Sie ein Programm, das den Benutzer auffordert, einen Benutzernamen und ein Passwort einzugeben. Überprüfen Sie, ob der Benutzername "admin" und das Passwort "password" sind. Geben Sie eine Erfolgsmeldung aus, wenn beides korrekt ist; andernfalls geben Sie eine Fehlermeldung aus.

## Weiter Syntax Möglichkeit für Selektion

Für jene die schnell unterwegs sind. 🏎️

### Theory

[Golang program that uses switch, multiple value cases](https://www.geeksforgeeks.org/golang-program-that-uses-switch-multiple-value-cases/)

### Aufgaben

#### 6. Wochentagsnamen

Schreibe ein Programm, das eine Zahl (1-7) vom Benutzer erhält und den entsprechenden Wochentagsnamen ausgibt.

#### 7. Monatsnamen

Schreibe ein Programm, das eine Zahl (1-12) vom Benutzer erhält und den entsprechenden Monatsnamen ausgibt.

#### 8. Notenbewertung

Schreibe ein Programm, das eine Zahl (0-100) als Punktzahl vom Benutzer erhält und die entsprechende Notenbewertung ausgibt:

- 90-100: A
- 80-89: B
- 70-79: C
- 60-69: D
- 0-59: F

## Lösungen

Versuche erst zuerst selbständig zu lösen bevor du hier dir Hilfe suchst. 👀

### 1. Gerade oder Ungerade

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Print("Geben Sie eine ganze Zahl ein: ")
    fmt.Scan(&num)

    if num%2 == 0 {
        fmt.Println("Die Zahl ist gerade.")
    } else {
        fmt.Println("Die Zahl ist ungerade.")
    }
}
```

### 2. Notenrechner

```go
package main

import "fmt"

func main() {
    var score int
    fmt.Print("Geben Sie die Note des Schülers ein: ")
    fmt.Scan(&score)

    if score >= 90 {
        fmt.Println("Note: A")
    } else if score >= 80 {
        fmt.Println("Note: B")
    } else if score >= 70 {
        fmt.Println("Note: C")
    } else if score >= 60 {
        fmt.Println("Note: D")
    } else {
        fmt.Println("Note: F")
    }
}
```

### 3. Altersklassifizierer

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Geben Sie das Alter der Person ein: ")
    fmt.Scan(&age)

    if age < 0 {
        fmt.Println("Ungültiges Alter")
    } else if age < 13 {
        fmt.Println("Kind")
    } else if age < 18 {
        fmt.Println("Jugendlicher")
    } else if age < 60 {
        fmt.Println("Erwachsener")
    } else {
        fmt.Println("Senior")
    }
}
```

### 4. Temperaturumrechner

```go
package main

import "fmt"

func main() {
    var temp float64
    var unit string

    fmt.Print("Geben Sie die Temperatur ein: ")
    fmt.Scan(&temp)
    fmt.Print("Geben Sie die Einheit ein (F oder C): ")
    fmt.Scan(&unit)

    if unit == "F" {
        temp = (temp - 32) * 5 / 9
        fmt.Printf("Umgerechnete Temperatur: %.2f°C\n", temp)
    } else if unit == "C" {
        temp = temp*9/5 + 32
        fmt.Printf("Umgerechnete Temperatur: %.2f°F\n", temp)
    } else {
        fmt.Println("Ungültige Einheit. Bitte geben Sie F oder C ein.")
    }
}
```

### 5. Benutzername und Passwortprüfer

```go
package main

import "fmt"

func main() {
    var username, password string

    fmt.Print("Geben Sie den Benutzernamen ein: ")
    fmt.Scan(&username)
    fmt.Print("Geben Sie das Passwort ein: ")
    fmt.Scan(&password)

    if username == "admin" && password == "password" {
        fmt.Println("Anmeldung erfolgreich!")
    } else {
        fmt.Println("Ungültiger Benutzername oder ungültiges Passwort. Bitte versuchen Sie es erneut.")
    }
}
```

#### 6. Wochentagsnamen

Lösung mit IF/ELSE

```go
package main

import "fmt"

func main() {
    var day int
    fmt.Print("Geben Sie eine Zahl (1-7) ein: ")
    fmt.Scan(&day)

    if day == 1 {
        fmt.Println("Montag")
    } else if day == 2 {
        fmt.Println("Dienstag")
    } else if day == 3 {
        fmt.Println("Mittwoch")
    } else if day == 4 {
        fmt.Println("Donnerstag")
    } else if day == 5 {
        fmt.Println("Freitag")
    } else if day == 6 {
        fmt.Println("Samstag")
    } else if day == 7 {
        fmt.Println("Sonntag")
    } else {
        fmt.Println("Ungültige Eingabe")
    }
}
```

Lösung mit SWITCH

```go
package main

import "fmt"

func main() {
    var day int
    fmt.Print("Geben Sie eine Zahl (1-7) ein: ")
    fmt.Scan(&day)

    switch day {
    case 1:
        fmt.Println("Montag")
    case 2:
        fmt.Println("Dienstag")
    case 3:
        fmt.Println("Mittwoch")
    case 4:
        fmt.Println("Donnerstag")
    case 5:
        fmt.Println("Freitag")
    case 6:
        fmt.Println("Samstag")
    case 7:
        fmt.Println("Sonntag")
    default:
        fmt.Println("Ungültige Eingabe")
    }
}
```

#### 7. Monatsnamen

Lösung mit IF/ELSE

```go
package main

import "fmt"

func main() {
    var month int
    fmt.Print("Geben Sie eine Zahl (1-12) ein: ")
    fmt.Scan(&month)

    if month == 1 {
        fmt.Println("Januar")
    } else if month == 2 {
        fmt.Println("Februar")
    } else if month == 3 {
        fmt.Println("März")
    } else if month == 4 {
        fmt.Println("April")
    } else if month == 5 {
        fmt.Println("Mai")
    } else if month == 6 {
        fmt.Println("Juni")
    } else if month == 7 {
        fmt.Println("Juli")
    } else if month == 8 {
        fmt.Println("August")
    } else if month == 9 {
        fmt.Println("September")
    } else if month == 10 {
        fmt.Println("Oktober")
    } else if month == 11 {
        fmt.Println("November")
    } else if month == 12 {
        fmt.Println("Dezember")
    } else {
        fmt.Println("Ungültige Eingabe")
    }
}
```

Lösung mit SWITCH

```go
package main

import "fmt"

func main() {
    var month int
    fmt.Print("Geben Sie eine Zahl (1-12) ein: ")
    fmt.Scan(&month)

    switch month {
    case 1:
        fmt.Println("Januar")
    case 2:
        fmt.Println("Februar")
    case 3:
        fmt.Println("März")
    case 4:
        fmt.Println("April")
    case 5:
        fmt.Println("Mai")
    case 6:
        fmt.Println("Juni")
    case 7:
        fmt.Println("Juli")
    case 8:
        fmt.Println("August")
    case 9:
        fmt.Println("September")
    case 10:
        fmt.Println("Oktober")
    case 11:
        fmt.Println("November")
    case 12:
        fmt.Println("Dezember")
    default:
        fmt.Println("Ungültige Eingabe")
    }
}
```

#### 8. Notenbewertung

Lösung mit IF/ELSE

```go
package main

import "fmt"

func main() {
    var score int
    fmt.Print("Geben Sie die Punktzahl (0-100) ein: ")
    fmt.Scan(&score)

    if score >= 90 && score <= 100 {
        fmt.Println("Note: A")
    } else if score >= 80 && score <= 89 {
        fmt.Println("Note: B")
    } else if score >= 70 && score <= 79 {
        fmt.Println("Note: C")
    } else if score >= 60 && score <= 69 {
        fmt.Println("Note: D")
    } else if score >= 0 && score <= 59 {
        fmt.Println("Note: F")
    } else {
        fmt.Println("Ungültige Eingabe")
    }
}
```

Lösung mit SWITCH

```go
package main

import "fmt"

func main() {
    var score int
    fmt.Print("Geben Sie die Punktzahl (0-100) ein: ")
    fmt.Scan(&score)

    switch {
    case score >= 90 && score <= 100:
        fmt.Println("Note: A")
    case score >= 80 && score <= 89:
        fmt.Println("Note: B")
    case score >= 70 && score <= 79:
        fmt.Println("Note: C")
    case score >= 60 && score <= 69:
        fmt.Println("Note: D")
    case score >= 0 && score <= 59:
        fmt.Println("Note: F")
    default:
        fmt.Println("Ungültige Eingabe")
    }
}
```
