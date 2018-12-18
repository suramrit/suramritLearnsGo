package main

import (
	"fmt"
)

//Some cool packages ..

func main() { //entry point to the program
	//Iodiomatic: Using patterns of speech in ways that they should be used...
	fmt.Println("Howdy Y'all!")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Guess we are at", i)
		}

	}

	foo()
	//exit flow from the program
}

func foo() {
	fmt.Println("Sup Matey!!")
}
