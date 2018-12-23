package main

import "fmt"

//further read: goinggo.net -- composition with go

type person struct {
	first string
	last  string
}

type student struct {
	person
	univ   string
	enroll int
	loan   int
}

func (p person) speak(s string) {
	fmt.Println(s)
}

func (st student) speak(s string) {
	fmt.Println(s)
}

//A value with fun speak attached speak() is of type 'human'
//A value can be of more than one type.
type human interface {
	speak(s string)
}

type universe interface {
	//empty interface -- that means that every type implements this interface
}

func plast(h human) {
	fmt.Println(h)
	//assertion - asserting something is of particular type.
	//switch on type
	switch h.(type) { //h.(type) -- special switch statement where you can switch on type.
	case person:
		fmt.Println("My first name is", h.(person).first)
	case student:
		fmt.Println("I have loan amounting to", h.(student).loan)
	}
}

//func (r receiver)<!attaches the function to the value of the type> identifier(param) (returns(s))

func main() {
	s1 := student{
		person: person{
			first: "suramrit",
			last:  "singh",
		},
		univ:   "U of Buffalo",
		enroll: 232523,
		loan:   18243,
	}
	p1 := person{
		first: "jackie",
		last:  "moss",
	}
	fmt.Println(s1, p1)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", p1)
	plast(s1) //polymorphism for plast due to s1 p1 implementing human interface
	plast(p1)
}
