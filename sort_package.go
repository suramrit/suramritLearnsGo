package main

import (
	"fmt"
	"sort" //sorting slices and user-defined types
)

type person struct {
	first string
	last  string
}

type pharma_emp struct {
	person
	liscence int
}

type pharm_data []pharma_emp

func (p pharm_data) Len() int {
	return len(p)
}
func (p pharm_data) Less(i, j int) bool {
	return (p)[i].liscence < (p)[j].liscence
}

func (p pharm_data) Swap(i, j int) {
	(p)[i], (p)[j] = (p)[j], (p)[i]
}

func main() {
	safeway := pharm_data{
		{
			person: person{
				first: "Jackie",
				last:  "Diane ",
			},
			liscence: 234635,
		},
		{
			person: person{
				first: "Char",
				last:  "parker",
			},
			liscence: 234633,
		},
	}
	//Using	sort.Slice
	//sort.Slice(safeway, func(i, j int) bool { return safeway[i].liscence < safeway[j].liscence })

	for _, v := range safeway {
		fmt.Println(v.person.first, "-----has liscence number-----", v.liscence)
	}

	
	// Interface interface :: from the docs --

	// 	A type, typically a collection, that satisfies sort.Interface can be sorted by the routines in this package.
	//  The methods require that the elements of the collection be enumerated by an integer index.
	// type Interface interface {
	//         // Len is the number of elements in the collection.
	//         Len() int
	//         // Less reports whether the element with
	//         // index i should sort before the element with index j.
	//         Less(i, j int) bool
	//         // Swap swaps the elements with indexes i and j.
	//         Swap(i, j int)
	// }

	// By making a collection of type Interface -- we can use functions in the sort package to sort the collections using sort.Sort method
	//Methods of Interface in sort package::
	/*
		- isSorted(data Interface )
		- Sort()
		- Stable()
		- Reverse(data Interface) Interface
	*/

	sort.Sort(safeway)

	fmt.Println("After Sorting ---- ")
	for _, v := range safeway {
		fmt.Println(v.person.first, "-----has liscence number-----", v.liscence)
	}

}
