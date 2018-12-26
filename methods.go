package main

import "fmt"

//Methods -- attached to a type --- here using the example of struct
// but it can be attached to any type

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



func (s *student) addloan(amt int) {
	s.loan = amt
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
	}
	fmt.Println(s1)
	s1.addloan(5000)
	fmt.Println(s1)
}
