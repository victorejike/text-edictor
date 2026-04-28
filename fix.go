package main

import (
	"fmt"
	"strings"
)

func FixPunc(text string) string {
	word := strings.Fields(text)
	result := []string{}

	for _, ch := range word {
		for len(ch) > 0 && strings.ContainsAny(ch[:1], ".,:;?!") {
			if len(result) > 0 {
				result[len(result)-1] += ch[:1]
			}
			ch = ch[1:]
		}
		if ch != "" {
			result = append(result, ch)
		}
	}
	return strings.Join(word, " ")
}

func main() {
	input := "a man   , .  .   has four  ,ford "
	output := FixPunc(input)

	fmt.Println("input", input)
	fmt.Println("output", output)
}
