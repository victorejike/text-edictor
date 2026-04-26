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

func MyFunction(fname string, age int) {
	fmt.Println("hello,", fname, "welcomeback!", "and do you know your ", age)
}

func main() {
	MyFunction("john", 18)
	MyFunction("victor", 20)
	MyFunction("ejike", 23)
	MyFunction("nmesomma", 12)
}
