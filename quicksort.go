package main

import (
	"fmt"
	"math/rand"
	"time"
)

//OS clock -- wall  (changes for sync) and monotonic (const.)

//var y = []int32{1234, 123, 0, -23, -15, 34, 12, 345, 25, 62, 26, 7, 234, 134, 5, 26, 23, -1, 234, 534, -1345, -134, 5345, 1345, 2345, 2346, 1, 316, 437, 13, 63, 6, 2436, 7, 7, 7, 7, 34, -234, 0, 2345, 6, 2346, 21, 4615, 6, 0}
var y = generateSlice(10000000) //10mill values took -- 20s 1st couple times~ then took 934 ms
//100mill- 27 sec
func main() {
	start := time.Now()
	s := qs(y)
	t := time.Now()
	elapse := t.Sub(start)
	//fmt.Printf("After sort:\n%v\n", s)
	fmt.Println("Sorting", len(s), "values took:", elapse, "seconds")
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func qs(x []int) []int {
	s := 0
	e := len(x) - 1
	if len(x) < 2 {
		return x
	}

	m := e - ((e - s) / 2)
	p := x[m]

	for s <= e {
		for x[s] < p {
			s++
		}
		for x[e] >= p && e > 0 {
			e--
		}
		if s <= e {
			x[s], x[e] = x[e], x[s]
			s++
			e--
		}
	}

	qs(x[0:s])
	qs(x[s:len(x)])

	return x
}
