package concurrency

import (
	"context"
	"fmt"
	"time"
)

//CONTEXT -- (Advanced Topic)

// package to pass request scoped values, cancelation signals, and deadlines across APIboundaries to all sub go routines
// involved in handling a single request

//Context type: struct carrying deadline, cancelation signal, request scoped variables across API bounderies which are safe for
// simultanous access through go routines.

//https://blog.golang.org/context
//https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8
//https://peter.bourgon.org/blog/2016/07/11/context.html

//ChannelContext Gives a simple example of how package context is used
func ChannelContext() {
	fmt.Println("Lets have a look at a bit more advanced topic - the context package")

	// gen generates integers in a separate sub-goroutine and
	// sends them to the returned channel.
	// The caller of gen needs to cancel the context
	// so as not to leak the internal goroutine started by gen.

	gen := func(ctx context.Context) <-chan int {
		//gen takes the context as an argument and returns a <-chan
		dst := make(chan int)
		n := 1

		fmt.Println("We are starting a goroutine that writes int vals onto a channel")
		time.Sleep(time.Second * 3)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("The subroutine got the signal that the context was canceled, so it will exit now")
					time.Sleep(time.Second * 3)
					return // returning not to leak the goroutine
				case dst <- n: //will block until a value is read from the dst channel
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println("Value read from channel", n)
		time.Sleep(time.Second * 1)
		if n == 15 {
			break
		}
	}
	fmt.Println("Context should be canceled now and the sub routine can exit safely")
}
