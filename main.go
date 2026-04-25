package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	// is to create a program that checks for file
	if len(os.Args) != 3 {
		fmt.Println("usage: for you to use this go run main.go <inputfile> <outputfile>")
		return
	}

	// now let me declear a variable for the inputfile and the outfile

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// now let me deceler a variable for our progam that is the data and for the error handling
	// using the os to read file to see what inside of our file

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error Reading this file please input the correct file")
		return
	}

	// now we have to create all our helper function that will handle all our changes

	text := string(data)

	text = transformer(text)


	err = os.WriteFile(outputFile, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error: cant write on this file check the file")
		return
	}
	fmt.Println("Your Program is completed!!")
}

func capitalized(text string)string{
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return result
}

func transformer(text string)string{
	token := strings.Fields(text)
	var result []string

	for i := 0; i < len(result); i++{
		token := token[i]

		if token == "(cap)" && len(result) > 0 {
			result[len(result)-1] = capitalized(result[len(result)-1])
			continue
		}

		if token == "(up)" && len(result) > 0 {
			
		}
	}
}