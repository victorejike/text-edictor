package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//the file handler the program that open the file and read the file

func main() {
	// using the os to pass the argument that if this not equal to 3
	if len(os.Args) != 3 {
		fmt.Println("usage: go run main.go <inputfile> <outputfile>")
		return
	}

	// declering a variable that will store the input and output

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// declera a variable  that will use to store our byte
	//and can handle error

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: Can't read this file please input the correct file")
		return
	}

	// let's  declera a variable for the helper functions where they will be called

	text := string(data)

	text = transforma(text)

	// now let us make it  to able to write on file which i will use the command to 0644 to givr permission to write on file

	err = os.WriteFile(outputFile, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error can't write on this file!!!!!!!!!")
		return
	}

	// this is for if the program ran succesful that should be the massage
	fmt.Println("your progamm has run succesfuly")

}

// this is the function that manage the capitalized word
func Capitalized(text string) string {
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return result
}

// the is where the transforamation happens
func transforma(text string) string {
	token := strings.Fields(text)
	var result []string

	for i := 0; i < len(result); i++ {
		token := token[i]
		//  this is the code that use the if statement to chack if the read file see somethings like this tis what it shuld

		if token == "(cap)" && len(result) > 0 {
			result[len(result)-1] = Capitalized(result[len(result)-1])
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

		// now we are going to checke for (hex) in our code if it finde any hex

		if token == "(hex)" && len(result) > 0 {
			val := result[len(result)-1]
			num, _ := strconv.ParseInt(val, 16, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}

		if token == "(bin)" && len(result) > 0 {
			val := result[len(result)-1]

			num, _ := strconv.ParseInt(val, 16, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}

		if token == "(cap," && i+1 < len(result){
			
		}

	}
}
