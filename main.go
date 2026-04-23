package main

import(
	"fmt"
	"os"
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
		result := strings.ToUpper(text) + strings.ToLower(text)
		return result
	}

	func transformer(text string)string  {
		 
	}
}