package main

import "fmt"

//Functions -- making thigns modular
//packages
//Everything in go is PASS BY VALUE!!

func main() {
	bar("yo wasssupp!")
	z := returner()
	fmt.Println(z)
	x, y := multi("jakie", "diane")
	fmt.Println(x)
	fmt.Println(y)
	sum(1, 2, 3, 45, 5, 6, 6, 7, 77)
	s := []int{1, 2, 3, 45, 5, 6, 6, 7, 77}
	sum()     //will run as sum() is variadic
	sum(s...) //pass slice as a simple list
}

// creating function
// func (r reciever) identifier(parameters) (returns(s)) { .. }

func foo() { //parameter
	fmt.Println("inside foo!! ")
}

func bar(s string) {
	fmt.Println(s)
}

func returner() string {
	defer foo() //deferred functions
	return "woo!"
}

//multiple return

func multi(fn, ln string) (string, bool) { //multiple returns need to be in commas
	a := fmt.Sprint(fn, " ", ln, ` said "hello!"`)
	b := true
	return a, b
}

//variadic parameters in functions
func sum(x ...int) {
	//Type:: []int .. pass individual items, which are stored in the function stack as a slice
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

//Modularising code

//Decoupling -- using defer
//deferred function is not run until the surrounding function executes and returns
