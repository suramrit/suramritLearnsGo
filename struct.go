package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type phrmagirl struct {
	person   // donot give a field name
	liscence int
	location string
}

//Embedded Structs -- enterin OOP in golang
// spec doc at --https://golang.org/ref/spec#Struct_types
type engineer struct {
	person
	field  string
	degree string
}

func main() {
	//Struct -- aggregate data type
	s := person{ //composite literal
		first: "Suramrit",
		last:  "Singh",
		age:   1000,
	}
	j := person{
		first: "jackie",
		last:  "diane", // will assign zero value to the age
	}
	fmt.Println(s.last, j, j.age)

	e1 := engineer{
		person: s,
		field:  "computer science",
		degree: "masters",
	}

	p1 := phrmagirl{ // If all the fields are different, can do without specifying them
		j,			 // Not a good practice !!! 
		318732,
		"Leavenworth",
	}
	fmt.Println(e1.first, p1) //inner type gets promoted to the outer type.. can specify namespace in case of a collision
	fmt.Println(p1.person)
	//Anon Structs -- used to avoid code pollution
	anon_struct := struct {
		name string
		age  int
	}{
		name: "suramrit",
		age:  22,
	}
	fmt.Println(anon_struct.age)
}
