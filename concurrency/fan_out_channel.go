package main 


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Fan out - Distribute work to different go routines and then fan-in 

 func main(){
 	c1:= make(chan int) // channel with the values that we need.. 
 	c2:= make(chan int)

 	go populate(c1) 

 	go fanOutIn_throttled(c1,c2)

 	i:=0
 	for v:= range c2 {
 		i++
 		fmt.Printf("Iteration:%d\n",i)
 		fmt.Println(v)
 	}

 	fmt.Println("EXIT")
 }


 func populate(c chan int){
 	for i := 0; i < 15; i++ {
 		c <- i
 	}
 	close(c)
 }

 //fan_out -- distribute items of a channel onto different go routines

 func fanOutIn(c1, c2 chan int) {
 	var wg sync.WaitGroup
 	for v:= range c1 {
 		wg.Add(1)
 		go func(v2 int){
 			c2 <- goDoSomeThing(v2) //we are giving the value to a single go routine, while collecting the result
 									// on a single channel -- fan out and then fan in... 
 			wg.Done()
 			}(v)
 	}
 	wg.Wait()
 	close(c2)
 }

func fanOutIn_throttled(c1,c2 chan int){ //Throttle the number of go routines that will be launched... 
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() { //Each routine will try can get some values to process until there are no more values in the channel to get
			count:=0
			for v := range c1 { // The range will be different for every go routine 
			count++
				func(v2 int) {
					c2 <- goDoSomeThing(v2)
				}(v)
			}
			fmt.Println("This routine ran",count,"times")
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func goDoSomeThing(v int) int{
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return (v)
}
