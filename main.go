package main

import(
	"fmt"
	"os"
	"strings"
	"strconv"
)
//the file handler the program that open the file and read the file

func main(){
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

func Capitalized(text string)string{
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return  result
}

func transforma(text string)string{
	token := strings.Fields(text)
	var result []string

	for i 
}