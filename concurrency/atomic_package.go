package main 

import (
	"fmt"
	"runtime" //Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines
	"sync"
	"sync/atomic" // provides low level atomic primitives for sync algorithms
)

// This program does the following -- 

// in main() we call 1000 Separate go routines,
// 		each routine updates the variable counter
// 		use atomic.AddInt64() and atomic.LoadInt64() to add delta to the int64 value and then read it without causing any race conditions 
// 		
// we also use sync.WaitGroup to make sure that main() does not exit until all the go routines are completed.


var wg sync.WaitGroup

func main() {
	//race conditions and resolutions --

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("goroutine\t", runtime.NumGoroutine()) //1

	var counter int64
	const g = 1000
	wg.Add(g)

	//Creating Race Condition --- 
	//Check when executing for race by: go run -race sample.go
	
	for i := 0; i < g; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			//runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter),counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(" after execution goroutine\t", runtime.NumGoroutine()) //1

	fmt.Println(counter)

}
