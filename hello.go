package main

import (
	"fmt"
)

var y string     // Declare with Identifier and TYPE as string, assign ZERO value of string to "z"
var z = "foobar" //entire program as scope .. limit global scope as best practice
var l = 42       // declaring l as a type int..
//Some cool packages ..
func main() {
	//entry point to the program
	//golang is Iodiomatic: Using patterns of speech in ways that they should be used...
	fmt.Println("Howdy Y'all!")
	fmt.Println(y)
	//Short declaration := both declare and assign value
	x := 42 //both declares variable and assigns a value
	fmt.Println(x)
	x = x + 3
	fmt.Printf("%T\n", x)
	//y := 2 + x
	y = "suramrit"
	//var vs := > := cannot be outside a function body
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	// for i := 0; i < 100; i++ {
	// 	if i%33 == 0 {
	// 		//Cant ignore the err return from the Println..
	// 		//Check the doc for return types.
	// 		n, _ := fmt.Println("Guess we are at", i, true, "didnt I say golang was cool!!")
	// 		fmt.Println(n)
	// 	}

	// }
	//foo()
	//Types in golang -- golang is all about type
	//Type can be printed using %T in fmt.Printf()
	x = 10051991
	fmt.Println(x)
	fmt.Printf("%T\n", x)       //type
	fmt.Printf("%b\n", x)       //binary
	fmt.Printf("%x\n", x)       //HEX
	fmt.Printf("%#x\n", x)      //0xHEX
	s := fmt.Sprintf("%b\n", x) // Binary as String
	fmt.Println(s + `suramrit   
					singh`) // `` for literal string

	//exit flow from the program
}
func foo() {
	fmt.Println("Sup Matey!!")
	fmt.Println("in foo", z)
}
