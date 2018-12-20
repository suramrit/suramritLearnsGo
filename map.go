package main

import "fmt"

func main() {
	// Map : Unordered list of Key-Value
	m := map[string]int{
		"suramrit": 28,
		"singh":    27,
	}
	fmt.Printf("%T\n", m)
	fmt.Println(m["not there"]) // Returns ZERO VALUE
	// , ok idiom  - used to check if a value exists
	v, ok := m["not there"]
	fmt.Println(v)
	fmt.Println(ok)

	if _, isok := m["not there"]; !isok {
		fmt.Println("Key does not exist !")
	}
	//ADD and for range for a map
	m["jaclyn"] = 27

	for k, v := range m {
		fmt.Println(k, v)
	}

	if _, ok := m["singh"]; ok {
		delete(m, "singh") //No errors if the key doesnt exist
		fmt.Println(ok)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
