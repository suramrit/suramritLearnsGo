package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//ChannelFanIn creates a simple fan-in channel that reads from two different channels
func ChannelFanIn() {

	even := make(chan int)
	odd := make(chan int)
	// FanIn is a concurrency pattern in which we use a single channel to collect values written over different channels in different goroutines
	fmt.Println("Lets have a look at how the concept of fan-in can be applied to simplify working with channels")
	time.Sleep(time.Second * 2)
	fan := make(chan int)

	go send(even, odd)
	go fanIn(even, odd, fan)

	for v := range fan {
		time.Sleep(time.Second * 2)
		fmt.Println("Read from fan - ", v)
	}

}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			time.Sleep(time.Second * 2)
			fmt.Println("Writing to even channel")
			even <- i
		} else {
			time.Sleep(time.Second * 2)
			fmt.Println("Writing to odd channel")
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// fan in -- Collect from different go routines into a single channel....

func fanIn(even, odd <-chan int, fan chan<- int) {

	var wg sync.WaitGroup

	wg.Add(2)

	//To streamline reading from all the channels we handle each read from a channel in a separate goroutine
	go func() {
		for v := range even {
			fan <- v
			time.Sleep(time.Second * 2)
			fmt.Println("Read from even channel and written onto fan")
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fan <- v
			time.Sleep(time.Second * 2)
			fmt.Println("Read from odd channel and written onto fan")
		}
		wg.Done()
	}()

	wg.Wait()
	close(fan)
}
