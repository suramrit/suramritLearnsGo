package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//ChannelFanOut creates a simple fan-out channel that writes to two different channels
func ChannelFanOut() {
	fmt.Println("In a fan out pattern, while writing onto a channel, we distribute that work to different go routines")
	time.Sleep(time.Second * 2)
	fmt.Println("We create two channels: c1,c2. We will write from c1 and read into c2")
	c1 := make(chan int) // channel with the values that we need..
	c2 := make(chan int)

	go populate(c1) // this go routine populates c1

	go fanOutInThrottled(c1, c2) //this goroutine spins up multiple go routines that try and read from c1 into c2

	i := 0
	for v := range c2 {
		//We use a for range loop to read from c2
		i++
		fmt.Printf("Iteration:%d\n", i)
		fmt.Println("value read: ", v)
	}
}

func populate(c chan int) {
	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 3)
		fmt.Print("In c1")
		fmt.Println(" writing - ", i)
		c <- i
	}
	close(c)
}

//fanOut -- distribute items of a channel onto different go routines

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- goDoSomeThing(v2) //we are giving the value to a single go routine, while collecting the result
			// on a single channel -- fan out and then fan in...
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func fanOutInThrottled(c1, c2 chan int) { //Throttle the number of go routines that will be launched...
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() { //Each routine will try can get some values to process until there are no more values in the channel to get
			time.Sleep(time.Second * 2)
			fmt.Println("This is a separate goroutine that is trying to read from c1 and write to c2")
			count := 0
			for v := range c1 { // The range will be different for every go routine
				count++
				time.Sleep(time.Second * 2)
				fmt.Println("This go-routine read:", v)

				func(v2 int) {
					c2 <- goDoSomeThing(v2)
				}(v)
			}
			fmt.Println("This routine ran", count, "times")
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func goDoSomeThing(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return (v)
}
