package main

import "fmt"

//Anonymous Functions
//Callback Functions
//functions expressions

var pckvar int //package level scope ...

func main() {
	//Anon...
	func() {
		fmt.Println("Anon hello")
	}() //trailing () to run the funciton code block
	//func expression -- assign function to a value
	v := func() {
		fmt.Println("expression")
	}
	//In golang -- functions are first class citizens
	v() 

	//First class -- so returning a function from a function
	x := f_return()
	x(`You are my Solskjaer!!`)
	f_return()(`" and the var said, "im a function!" "s`)

	//CallBacks -- passing func as an argument
	//Functional Programming!
	fmt.Println(even_sum(sum, 12, 3, 4, 5, 1, 3, 5, 61, 6, 61, 34, 6, 1346))

	//Closure: Close the scope of a variable to limit it ..
	//scope can be a. package level -- x, b. function level -- x, c. block level -- 'i' in the loops
	in := incrementor()
	fmt.Println(in()) //100
	fmt.Println(in()) //101 !! -- Think of this way..
	// You defined 'x' incrementor, which defines its scope
	// in:=incrementor will call the init of 'x'
	// any new time the function in() is called, it will increment the value that was defined in increment() intially...
	new_in := incrementor()
	fmt.Println(new_in()) //100 again .. this time a new variable will be declared and stored...
}

func incrementor() func() int {
	var x int
	x = 99
	return func() int {
		x++
		return x
	}
}

func even_sum(f func(x ...int) int, y ...int) int {
	var eve []int
	for i, _ := range y {
		if y[i]%2 == 0 {
			eve = append(eve, y[i])
		}
	}
	return f(eve...)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func f_return() func(s string) {
	s := func(s string) {
		fmt.Println("Is this:", s, "what was needed")
	}
	return s
}
