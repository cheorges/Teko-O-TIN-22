package main

import (
	"fmt"
	"strings"
	"time"
)

func replacer() {
	broken := "G# r#cks!"
	// 1er replacer welcher # durch o ersetzt
	dashReplacer := strings.NewReplacer("#", "o")
	fixed := dashReplacer.Replace(broken)
	fmt.Println(fixed)

	// 2er replacer welcher ! durch yeahh!!! ersetzt
	screamReplacer := strings.NewReplacer("!", " yeahh!!!")
	greetingFixed := screamReplacer.Replace(broken)
	fmt.Println(greetingFixed)

}

func maain() {
	var now time.Time = time.Now() // Function call
	var year int = now.Year()      // Method call
	fmt.Println(year)              // Function call
}
