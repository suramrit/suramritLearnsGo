package main

import "fmt"

//a look at Array, Slice, Struct and Map in golang

func main() {

	//Slices::
	y := []int{4, 5, 6, 7, 8} //composite literal
	//SLICE allows to group together values of same type
	fmt.Printf("%v\t%T\t%v\t%v\n", y, y, len(y), cap(y))

	//range clause
	for i, v := range y {
		fmt.Println(i, v)
	}
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}

	//Sicing a slice -- getting sub slices
	fmt.Println(y[1:4]) // SLICE[start position : upto position]

	//append to slice
	//append() special builtin function in package "builtin" --
	y = append(y, 9, 10, 11)
	fmt.Println(cap(y)) // returns 10: cap is the max length the slice may assume
	fmt.Println(y[:])
	// y = append(y, 12, 13, 14, 15)

	// fmt.Println(cap(y)) // return 20: cap is amortized to 2*cap when data exceeds the capacity
	fmt.Println(y[:])
	y = append(y, y...) //SLICE... -- list out elements of the Slice
	fmt.Println(y[:])

	//EFFECTIVE DELETING FROM SLICE -- USE APPEND!!!!!!
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	//Making tasty Slices (make function)--
	//make([]T, length, capacity)
	//slice built on top array but dynamic
	//make() - can give initial size of the slice that we want to create, saving runtime for dynamic allocation
	z := make([]int, 10, 100)
	fmt.Println(z) //empty slice len(10) cap(100)
	// fmt.Println(z[11]) -- index out of range
	z = append(z, 1234)
	fmt.Println(z, cap(z))

	//Multi dimensional Slice
	md := []string{"Name", "Place of Birth", "Birth Day", "Month"}
	fmt.Println(md)
	data1 := []string{"Suramrit", "Jammu", "5", "October"}
	fmt.Println(data1)
	data2 := []string{"Plain", "27", "October"}
	fmt.Println(data2)

	multi := [][]string{md, data1, data2} //valid -- even if data1 data2 have diff dimensionality
	fmt.Println(multi)
	//fmt.Println(multi[2][3])             //index out of bounds
	for i, v := range multi {
		fmt.Println(i, v)
	}
}
