package main 

import(
"fmt"
"sync")


func main(){

	even:= make(chan int)
	odd:= make (chan int)
	fan:= make(chan int)

	go send(even, odd)
	go fan_in(even, odd, fan)

	for v:= range fan{
		fmt.Println(v)
	}

}

func send(even, odd chan<- int){
	for i := 0; i < 10; i++ {
		if i%2==0 {
			even<-i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}


// fan in -- Collect from different go routines into a single channel.... 


func fan_in(even, odd <-chan int, fan chan<- int) {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for v:= range even{
			fan <- v
		}
		wg.Done()
	}()

go func() {
	for v:= range odd{
		fan <- v
	}
	wg.Done()
}()

wg.Wait()
close(fan)
}