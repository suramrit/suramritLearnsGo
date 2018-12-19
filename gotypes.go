package main

import "fmt"

//GO IS ALL ABOUT TYPE

//Cerating user defined TYPE
//Conversions and casting in go

var a bool
var x = 42
var y = "suramrit"
var z = true

type suramrit bool

var jaclyn suramrit
var moss bool

func main() {
	a := true // This is not an Error
	fmt.Println(a)
	var s suramrit = false
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	jaclyn = true
	//a = s // Cannot use as TYPE do not comply
	// Incorrect : fmt.Printf("%T",suramrit)
	//Conversions  --- If the underlying TYPE are same, we can do CONVERSIONS
	s = suramrit(a) // converting "a" to TYPE "suramrit"
	fmt.Println(s)
	t := fmt.Sprintf("%v%v%v", x, y, z) //
	fmt.Println(t)
	fmt.Println(jaclyn)
	fmt.Printf("Jaclyn is %T\n", jaclyn)

	moss := bool(jaclyn)
	fmt.Println(moss)
	fmt.Printf("%T\n", moss)

	// !! Not called "casting" in golang
}
