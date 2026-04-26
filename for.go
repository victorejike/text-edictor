// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fruite := [3]string{"apple", "mango", "orange"}

//		for i, j := range fruite {
//			fmt.Printf("%v \t %v\n", i, j)
//		}
//	}
package main

import (
	"fmt"
)

func MyFunction(fname string) {
	fmt.Println("hello,", fname, "welcomeback!")
}

func main() {
	MyFunction("john")
	MyFunction("victor")
	MyFunction("ejike")
	MyFunction("nmesomma")
}
