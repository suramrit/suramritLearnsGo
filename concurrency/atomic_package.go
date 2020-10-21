package concurrency

import (
	"fmt"
	"runtime"     // Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines
	"sync/atomic" // Atomic provides low level atomic primitives for sync algorithms
)

// This program does the following --

// AtomicPackage calls 1000 Separate go routines,
// 		each routine updates the variable counter
// 		use atomic.AddInt64() and atomic.LoadInt64() to add delta to the int64 value and then read it without causing any race conditions
//
// we also use sync.WaitGroup to make sure that AtomicPackage() does not exit until all the go routines are completed.
func AtomicPackage() {
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
			// Alternative to using mutex primitive
			grNum := runtime.NumGoroutine()
			atomic.AddInt64(&counter, 1)
			//runtime.Gosched()
			fmt.Println("For goroutine number - ", grNum, "counter value is at:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Things to note:")
	fmt.Println("- We dont necessarily get 1000 goroutines")
	fmt.Println("- Print is not sequential as they could have been scheduled differently")
	fmt.Println(" after execution goroutine\t", runtime.NumGoroutine()) //1
	fmt.Println(counter)

}
