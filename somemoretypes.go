package main

import (
	"fmt"
	"runtime" // package to check runtime arch of the system
)

const x string = "suramrit"


func main() {
	//bool in action
	x := true
	fmt.Println(x)
	//numerical
	// go by default int and float64
	a := 99
	b := 22.34
	fmt.Println(a)
	fmt.Printf("%T\n", b)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	//String
	s := `"hey yo"`
	bytestring := []byte(s) // String in implementation is a Byte Sequence
	fmt.Printf("%#x\n", bytestring)
	fmt.Printf("%T\n", bytestring)

	for i, v := range s { // top explore further in control 
		fmt.Printf("at index %d, hex %#x\n", i, v)
	}

	//Constants 
	const(
		a1 = 24
		b1 = 55.4
		c1 = false
		d1 = iota	//Useful with const 
	)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1) // d1 == 4 

	//Bit Shifting 
	a = 4
	fmt.Println(a<<4)
	//iota can be used for bit shifting 
	const(
		_ = iota // couldnt use :=
		//kb:= 1024
		kb= 1 << (iota*10)
		mb= 1 << (iota*10)
		gb= 1 << (iota*10)
	)

	fmt.Printf("GB in decimal %d, binary %b\n", gb, gb)


}
