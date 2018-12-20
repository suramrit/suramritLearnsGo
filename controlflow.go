package main

import "fmt"

//Covering almost everything about Control flow in golang
// Good article - Vladimir medoum.com - hacking bit shifting with golang
// Good references - gobyexample.com
func main() {
	//for with caluse -- no WHILE in golang!
	for i := 0; i < 5; i++ {
		fmt.Println("Are at", i)
	}
	//print unicode chars
	for i := 33; i <= 122; i++ {
		fmt.Printf("Index:%d\tUnicode:%#U\t \n", i, i)
	}

	//nested
	for i := 0; i < 2; i++ {
		fmt.Println("Knock it off!")
		for j := 0; j < 1; j++ {
			if i%2 == 0 {
				continue
			}

			fmt.Println("Whatcha doing?")
		}
	}

	//for with single condition
	a := 0
	b := 3
	for a < b { //equivalent to a while block
		a++
		fmt.Println(a)
	}
	//for with no condition
	for {
		fmt.Println("break!")
		break
	}
	//print all first 10 evens
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 == 0 {
			fmt.Println(x)
		}
	}

	//TODO for with range -- iterates over enteries of array slice, etc

	//conditionals
	if s := 1; s != 1 {
		if !(true != true) {
			fmt.Println("voila")
		}
	} else if s == 1 {
		fmt.Println("value is one")
	}
	//fmt.Println(s) -- will not work.. scope was limited

	//switch statements
	switch {
	case true:
		fmt.Println("Prints")
		fallthrough //fallthrough is a keyword -- causes fallthrough to the next case even if its false.
		//retrict use of fallthrough
	case false:
		fmt.Println("Prints")
		fallthrough
	case true:
		fmt.Println("Doesnt print") // no default fall through for switch for golang
	}

	//switch to a value
	n := "suramrit"
	t := "jaclyn"
	switch t {
	case "jaclyn":
		fmt.Println("This is Jaclyn, how may I assist you?")
	case n:
	}
	hh := 'k'
	fmt.Printf("%c\t%T\n", hh, hh)
	//switch with expressions
	switch 's' { // can be expression
	case 's', 'j', hh: // can be expression, or expression list
		fmt.Println("Found the match")
	}

	// for i := 0; i < 29; i++ {
	// 	fmt.Println(1991 + i)
	// }
	// bd := 1991
	// for {
	// 	fmt.Println(bd)
	// 	if bd == 2019 {
	// 		break
	// 	}
	// 	bd++
	// }
	for i := 10; i < 100; i++ {
		fmt.Println(i % 4)
	}
	//boolean simple examples
	// true && false
	// true || true
	// !trues

}
