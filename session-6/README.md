# Unterrichtsmaterial: Arrays, Slices und Maps in Go

## Lektion 1: Einführung in Arrays

### Theorie
- **Definition**: Ein Array ist ein fester, homogener Datentyp, der eine Sammlung von Elementen desselben Typs enthält. Die Größe des Arrays muss bei der Deklaration festgelegt werden und kann nicht verändert werden.
- **Deklaration und Initialisierung**:
    ```go
    var arr [5]int // deklariert ein Array mit 5 ganzen Zahlen
    arr[0] = 10    // Initialisiert das erste Element mit dem Wert 10

    // Oder direkt initialisiert:
    arr := [5]int{1, 2, 3, 4, 5}
    ```

- **Zugriff auf Elemente**: Elemente eines Arrays werden über ihren Index angesprochen, wobei der Index bei 0 beginnt.
    ```go
    fmt.Println(arr[0]) // Ausgabe: 10
    ```

### Aufgabe 1
Deklariere ein Array von 7 Zeichen und initialisiere es mit den Buchstaben "G", "o", "l", "a", "n", "g", "!".
```go
arr := [7]rune{'G', 'o', 'l', 'a', 'n', 'g', '!'}
```

### Aufgabe 2
Erstelle ein Array von 5 Fließkommazahlen und weise den ersten beiden Elementen die Werte 1.5 und 2.7 zu.
```go
arr := [5]float64{1.5, 2.7}
```

## Lektion 2: Iteration über Arrays

### Theorie
- **for-Schleife**: Eine traditionelle for-Schleife kann verwendet werden, um über die Elemente eines Arrays zu iterieren.
    ```go
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    ```
- **Range-Schleifen**: Eine range-Schleife bietet eine einfachere Möglichkeit, über die Elemente eines Arrays zu iterieren. Dabei werden sowohl der Index als auch der Wert der Elemente zur Verfügung gestellt.
    ```go
    for index, value := range arr {
        fmt.Printf("Index: %d, Wert: %d\n", index, value)
    }
    ```

### Aufgabe 1
Iteriere über das Array aus der letzten Aufgabe und gib jeden Buchstaben aus.
#### Lösung
```go
for i := 0; i < len(arr); i++ {
    fmt.Println(string(arr[i]))
}

for _, value := range arr {
    fmt.Println(string(value))
}
```

### Aufgabe 2
Iteriere über das Array von Fließkommazahlen und gib die Werte und ihre Quadrate aus.
#### Lösung
```go
for _, value := range arr {
    fmt.Printf("Wert: %.2f, Quadrat: %.2f\n", value, value*value)
}
```

## Lektion 3: Einführung in Slices

### Theorie
- **Definition**: Ein Slice ist eine flexible und dynamische Datenstruktur, die ein Array darstellt. Die Länge eines Slices kann während der Laufzeit geändert werden.
- **Deklaration und Initialisierung**:
    ```go
    var s []int             // deklariert ein Slice ohne Initialisierung
    s = make([]int, 5)      // Initialisiert ein Slice mit 5 Elementen

    // Direkt initialisiert:
    s := []int{1, 2, 3, 4, 5}
    ```
- **Zugriff auf Elemente**: Ähnlich wie bei Arrays können Elemente über ihren Index angesprochen werden.
    ```go
    fmt.Println(s[1]) // Ausgabe: 2
    ```
- **Länge und Kapazität**: Die Länge eines Slices gibt die Anzahl der enthaltenen Elemente an, während die Kapazität die maximal mögliche Anzahl von Elementen angibt.
    ```go
    fmt.Printf("Länge: %d, Kapazität: %d\n", len(s), cap(s))
    ```

### Aufgabe 1
Erzeuge ein Slice von 5 Zeichen und initialisiere es mit Buchstaben "A" bis "E".
#### Lösung
```go
s := []rune{'A', 'B', 'C', 'D', 'E'}
```

### Aufgabe 2
Erstelle ein Slice von 3 ganzen Zahlen und füge die Werte 4 und 5 hinzu.
#### Lösung
```go
s := []int{1, 2, 3}
s = append(s, 4, 5)
```

## Lektion 4: Slices und Arrays

### Theorie
- **Array in Slice umwandeln**: Ein Slice kann aus einem Array erstellt werden, wobei das Slice auf den gleichen Speicherbereich verweist.
    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    s := arr[:] // Erstellt ein Slice, das auf das gesamte Array zeigt
    ```
- **Slices erweitern**: Slices können dynamisch erweitert werden, indem neue Elemente angehängt werden. Dabei kann die Kapazität des Slices bei Bedarf automatisch vergrößert werden.
    ```go
    s := []int{1, 2, 3}
    s = append(s, 4, 5) // Fügt die Werte 4 und 5 am Ende des Slices hinzu
    ```

### Aufgabe 1
Erweitere das Slice aus der letzten Aufgabe um die Buchstaben "F" bis "H".
#### Lösung
```go
s = append(s, 'F', 'G', 'H')
```

### Aufgabe 2
Erzeuge aus einem Array von ganzen Zahlen ein Slice, und füge weitere Elemente hinzu.
#### Lösung
```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[:]
s = append(s, 60, 70)
```

## Lektion 5: Slices: Unter-Slices erstellen

### Theorie
- **Unter-Slices**: Ein Unter-Slice kann durch Angabe von Bereichsgrenzen erstellt werden. Unter-Slices teilen sich den gleichen Speicherbereich wie das Original-Slice.
    ```go
    arr := []int{1, 2, 3, 4, 5, 6}
    sub := arr[1:4] // Enthält die Elemente von Index 1 bis 3
    ```

### Aufgabe 1
Erzeuge ein Unter-Slice des Slices aus Lektion 4, das die Buchstaben "C" bis "E" enthält.
#### Lösung
```go
sub := s[2:5]
```

### Aufgabe 2
Erstelle ein Unter-Slice eines Slices von Fließkommazahlen von Index 1 bis einschließlich 3.
#### Lösung
```go
f := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
sub := f[1:4]
```

## Lektion 6: Einführung in Maps

### Theorie
- **Definition**: Eine Map ist eine Sammlung von Schlüssel-Wert-Paaren, die eine schnelle Zuordnung und Abfrage ermöglichen. Die Schlüssel müssen eindeutig und von einem geeigneten Typ sein.
- **Deklaration und Initialisierung**:
    ```go
    var m map[string]int
    m = make(map[string]int)

    // Oder direkt initialisiert:
    m := map[string]int{
        "eins": 1,
        "zwei": 2,
    }
    ```
- **Zugriff auf Werte**: Werte können über ihren Schlüssel abgerufen und geändert werden.
    ```go
    m["eins"] = 1     // Weist den Wert 1 dem Schlüssel "eins" zu
    value := m["zwei"] // Ruft den Wert zum Schlüssel "zwei" ab
    ```
- **Überprüfen auf Vorhandensein eines Schlüssels**:
    ```go
    value, exists := m["drei"]
    if exists {
        fmt.Println("Schlüssel 'drei' existiert mit Wert:", value)
    } else {
        fmt.Println("Schlüssel 'drei' existiert nicht.")
    }
    ```

### Aufgabe 1
Erstelle eine Map, die die Anzahl der einzelnen Buchstaben in einem Text zählt.
#### Lösung
```go
text := "golang"
letterCount := make(map[rune]int)
for _, letter := range text {
    letterCount[letter]++
}
```

### Aufgabe 2
Erstelle eine Map von Studentennamen als Schlüssel und deren Alter als Wert.
#### Lösung
```go
students := map[string]int{
    "Alice": 21,
    "Bob": 23,
    "Charlie": 20,
}
```

## Lektion 7: Iteration über Maps

### Theorie
- **Range-Schleifen**: Über eine Map kann iteriert werden, ähnlich wie bei Arrays und Slices. Dabei werden sowohl der Schlüssel als auch der Wert zur Verfügung gestellt.
    ```go
    for key, value := range m {
        fmt.Printf("Schlüssel: %s, Wert: %d\n", key, value)
    }
    ```

### Aufgabe 1
Iteriere über die Map aus der letzten Aufgabe und gib jeden Buchstaben und seine Häufigkeit aus.
#### Lösung
```go
for key, value := range letterCount {
    fmt.Printf("Buchstabe: %c, Häufigkeit: %d\n", key, value)
}
```

### Aufgabe 2
Iteriere über die Map von Studentennamen und Alter und gib die Informationen aus.
#### Lösung
```go
for name, age := range students {
    fmt.Printf("Name: %s, Alter: %d\n", name, age)
}
```

## Lektion 8: Maps, Arrays und Slices kombinieren

### Theorie
- **Maps von Arrays/Slices**: Maps können Arrays oder Slices als Werte verwenden, was eine flexible Datenstruktur für unterschiedliche Datentypen und -größen ermöglicht.
    ```go
    m := make(map[string][]int)
    m["zahlen"] = []int{1, 2, 3}
    ```
- **Erweiterbare Datenstrukturen**: Der Vorteil der Kombination von Maps und Slices liegt in der dynamischen Anpassung von Datenstrukturen bei gleichzeitiger effizienter Zuordnung und Abfrage.
    ```go
    m["zahlen"] = append(m["zahlen"], 4, 5)
    ```

### Aufgabe 1
Erstelle eine Map, die Studentennamen als Schlüssel und deren Noten als Slice von int enthält.
#### Lösung
```go
students := make(map[string][]int)
students["Alice"] = []int{90, 85, 88}
students["Bob"] = []int{70, 75, 72}
```

### Aufgabe 2
Füge der Map aus Aufgabe 1 weitere Noten für die Studenten hinzu.
#### Lösung
```go
students["Alice"] = append(students["Alice"], 92)
students["Bob"] = append(students["Bob"], 68)
```