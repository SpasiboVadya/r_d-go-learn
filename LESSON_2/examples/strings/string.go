package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	greeting := string("Привіт світ!")
	greeting = strings.ReplaceAll(greeting, "т", "д")
	fmt.Println(greeting)
	greetingRune := []rune(greeting)
	for i, char := range greetingRune {
		fmt.Printf("%d: %c\n", i, char)
	}

	fmt.Printf("%v, len: %d", []byte(greeting), utf8.RuneCountInString(greeting))
}
