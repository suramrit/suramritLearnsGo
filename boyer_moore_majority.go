package main
import "fmt"
func main() {
	var team = []int{0, 3, 4, 5, 6, 1, 3, 4, 56,3, 7, 3,1, 3,2, 3,34,3, 3, 3, 3, 3}

	// simple majority -- count(e) > int(n/2)

	candidate := team[0]

	count := 0

	for _, v := range team {
		if count == 0{
			candidate = v
		} 
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	// The candidate after this pass is the majority value. 
	fmt.Println(count)
	fmt.Println(candidate)
	fmt.Println(len(team))
	c_count:=0
	for _,v := range team {
		if v == candidate{c_count++}
	}
	fmt.Println(c_count)
	// The validity of the majority can be confirmed by another pass over the A

}
