package main

import (
	"fmt"
	"go/token"
	"os"
	"strconv"
	"strings"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage: run main.go <inputFile> <outputFile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading this file please try intputfile dear")
		return
	}

	text := strings(data)

	text = transformer(text)
	text = fixacticle(text)
	test = fixpuncation(text)
	test = fixquotes(text)
 
	err = os.WriteFile(outputFile,[]byte(data), 0644)
	if err != nil {
		fmt.Println("Error wrinting to file ")
		return
	}

	fmt.Println("Your Program just ran sucssfully")


	func Capitalize(text string) string {
		result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
		return result
	}

	func transformer(text string)string  {
		 token := strings.Fields(text)
		 var result []string

		 for i := 0; i < len(token); i++{
			token := token[i]

			if token == "(cap)" && len(result) > 0 {
				result[len(result)-1] = Capitalize(result[len(result)-1])
				continue
			}

			if token == "(up)" && len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
				continue
			}

			if token == "(low)" && len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
				continue
			}

			if token == "(hex)" && len(result) > 0 {
				val := result[len(result)-1]
				num, _ := strconv.ParseInt(val, 16, 64)
				result[len(result)] = strconv.FormatInt(num, 10)
				continue
			}
		 }
	}
}