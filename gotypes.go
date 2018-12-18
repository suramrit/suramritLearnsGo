package main

import "fmt"

//GO IS ALL ABOUT TYPE

//Cerating user defined TYPE
//Conversions and casting in go

var a bool

type suramrit bool

func main() {
	a := true // This is not an Error
	fmt.Println(a)
	var s suramrit = false
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//a = s // Cannot use as TYPE do not comply
	// Incorrect : fmt.Printf("%T",suramrit)
	//Conversions  --- If the underlying TYPE are same, we can do CONVERSIONS
	s = suramrit(a) // converting "a" to TYPE "suramrit"
	fmt.Println(s)
	// !! Not called "casting" in golang
}
