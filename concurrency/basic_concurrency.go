package concurrency

import (
	"fmt"
	"runtime" //Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines
	"sync"
)

//Concurrency in golang --
//Go desgined to utilised multi core processors..
//Write code which handles different things parallely
//Rob Pike (Co-creator golang)
//		- Concurrency is dealing with a lot of things at once, parallelism is doing a lot of things at once
//		- Separate ideas. Concurrency - structure, Parallelism - execution
//		- Paper: Tony Hoare - "communicating with sequential processes" 1978
//		- Adding concurrency can make things efficient, even if it requires more work to be done
// 		- But parallelism is only achieved when the concurrent pieces are allowed to work at the same time

// This program does the following --

// f1() and f2() concurrently print numbers upto 9 s

// in BasicConcurrency() we call 1000 Separate go routines,
// 		each routine updates the varaible counter
// 		if no mutex used, race condition causes counter to be incorrect
// 		mutex use makes sure theres no race condition
// 		we also use sync.WaitGroup to make sure that BasicConcurrency() does not exit until all the go routines are completed.

var wg sync.WaitGroup

// BasicConcurrency inits 1000 separate goroutines that update the same variable which is under a mutex
func BasicConcurrency() {
	//Print Sys Config
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("Arch\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	//Print the number of current goroutine
	fmt.Println("goroutine\t", runtime.NumGoroutine()) //1

	// add 2 the the our count in the wait group
	wg.Add(2)
	//f1 will count till 9
	go f1()
	//or --
	//var wg sync.WaitGroup
	//f2 will count till 9                             //Not printed -- launches a new go routine
	go f2()                                            //Not printed -- not printed.. becuase the program exited... causing the goroutines to be incomplete....
	fmt.Println("goroutine\t", runtime.NumGoroutine()) //3

	//Synchronization primitives: mutex, waitgroups; sync package

	//race conditions and resolutions --

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("goroutine\t", runtime.NumGoroutine()) //Can be anything between 1-3 depending on how f1 f2 run

	counter := 0
	const g = 1000
	wg.Add(g)
	//MUTEX!!
	var mute sync.Mutex
	//Creating Race Condition ---
	//Check when executing for race by: go run -race sample.go
	for i := 0; i < g; i++ {
		go func() {
			//NEED MUTEX!!!!
			mute.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mute.Unlock()
			wg.Done()
		}()
	}
	//Wait blocks until all go routines are complete
	wg.Wait()

	fmt.Println(" after execution goroutine\t", runtime.NumGoroutine()) //1

	fmt.Println(counter)

}

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println("f1:", i)
	}
	wg.Done()
}

func f2() {
	for i := 0; i < 10; i++ {
		fmt.Println("f2:", i)
	}
	wg.Done()
}
