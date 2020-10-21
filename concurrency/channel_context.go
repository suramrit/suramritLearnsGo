package main 

import("fmt"
	"context"
	)

//CONTEXT -- (Advanced Topic)

// package to pass request scoped values, cancelation signals, and deadlines across APIboundaries to all sub go routines 
// involved in. handling a single request 

//Context type: struct carrying deadline, cancelation signal, request scoped variables across API bounderies which are safe for 
// simultanous access through go routines. 


//https://blog.golang.org/context 
//https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8 
//https://peter.bourgon.org/blog/2016/07/11/context.html 

// Touching just the surface for now, will cover in detial as I gain more familiarity with the topic. 

func main() {
	fmt.Println("ENTRY")

	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 55 {
			break
		}
	}


}