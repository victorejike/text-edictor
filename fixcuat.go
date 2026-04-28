package main

import (
	"fmt"
	"strings"
)

func fixcuat(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiouhAEIOUH", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}

func main() {
	input := "a appple a elephant a ant  a House   a unblera "

	output := fixcuat(input)

	fmt.Println("input", input)
	fmt.Println("output", output)
}
