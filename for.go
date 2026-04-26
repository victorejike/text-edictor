package main

import (
	"fmt"
)

func main() {
	fruite := [3]string{"apple", "mango", "orange"}

	for i, j := range fruite {
		fmt.Printf("%v %s\n", i, j)
	}
}
