package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).last = "singh"
	//p.last ="singh" -- also works
}

func main() {
	//EVERYTHING IN golang IS PASS BY VALUE!!!
	//working on pointers after almost 3 years.... this is going to be fun! :D
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)
	var b *int = &a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("%T\n", &b)
	*b = 9
	fmt.Println(a)
	foo(a)
	fmt.Println(a)
	bar(&a)
	fmt.Println(a)
	//Pointers useful for communicating large chunks of data...
	c := circle{10}
	// info(c) will not work as long as area() has *circle as the receiver
	info(&c)

	p := person{
		first: "suramrit",
		last:  "",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}

//Method Sets -- set of methods attached to a type
// In the method set for an interface...
// if the receiver of a method is of type T -- then it can have *T or T as value
// if the receiver of a method is of type *T -- then it can only have *T as a valu
type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

type circle struct {
	radius int
}

func (c *circle) area() {
	a := c.radius * c.radius * 3 // its a joke -- chill!
	fmt.Println("area:", a)
}

func foo(x int) {
	fmt.Println(x)
	x++
	fmt.Println(x)
}

func bar(x *int) {
	fmt.Println(*x)
	*x++
	fmt.Println(*x)
}
