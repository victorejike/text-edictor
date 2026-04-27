package main

import (
	"fmt"
	"strings"
)

func Fixarticle(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiouh", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}

func main() {
	input := "a apple a house an apple a eroplain"
	output := Fixarticle(input)

	fmt.Println("input", input)
	fmt.Println("output", output)

}
