package main

import "fmt"

//a look at Array, Slice, Struct and Map in golang

func main() {

	//Aggregates -- values of same type

	//Arrays::
	var x [5]int //size needs to be specified
	fmt.Println(x)
	x[3] = 42
	// There are major differences between the ways arrays work in Go and C. In Go,
	// Arrays are values. Assigning one array to another copies all the elements.
	// In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
	// The size of an array is part of its type. The types [10]int and [20]int are distinct.

	//Using arrays is not idiomatic to golang!
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("%T\n", x) // Type: [5]int

	//range clause
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
