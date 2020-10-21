package concurrency

import (
	"fmt"
	"time"
)

// Channels creates difffernt kinds of channels supported in golang and observes their behavior
func Channels() {
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Creating a simple channel")
	c := make(chan int)
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Channels Block!")
	//c <- 32
	//This will block..
	//until send and recieve at the same time...
	//or  buffered channel
	go func() {
		fmt.Println("This is a separate goroutine where we write to the channel. This go routine will block until someone reads from the channel")
		c <- 332 // Write to channel .. go routine is blocked until channel is read
		time.Sleep(time.Millisecond * 2000)
		fmt.Println("Phew! Someone finally read from my channel!")
	}()
	time.Sleep(time.Millisecond * 5000)
	fmt.Println("Now back in main goroutine and reading will cause the other routine to unblock") // reads from the channel... once read, goroutine can unblock
	//let him be released!
	<-c
	// Directional Channels - only send to or recieve from channel
	time.Sleep(time.Millisecond * 5000)
	fmt.Printf("Channels are strongly typed. This is the type of a basic channel -  %T\t\n", c)

	onlySendChannel := make(chan<- int) // Can only send to the channel
	recvOnly := make(<-chan int)        // Can only recieve from the channel
	time.Sleep(time.Millisecond * 4000)
	fmt.Printf("Channel can also be of type send only like this channels type is - %T\t\n", onlySendChannel)
	time.Sleep(time.Millisecond * 2000)
	fmt.Printf("or recieve only - %T\t\n", recvOnly)
	// recv_only <- 32 //error!
	go func() {
		onlySendChannel <- 223
	}()
	// fmt.Println(<- only_send_channel) //error!

	// this go routine takes the channel c and add values to it
	time.Sleep(time.Millisecond * 5000)
	fmt.Println(" We can use a for-range loop to read from channel ")
	go sender(c)

	//Now how to read from this channel 'c'  ?? use a for-range loop!

	// Ranging over a channel -- blocks until the channel is closed
	// No error when ranging over a closed channel
	// Error/blocks with deadlock when channel is left unclosed/open
	for v := range c {
		time.Sleep(time.Millisecond * 2000)
		fmt.Println("reading from channel c in the loop", v)
		// will result in deadlock until the channel is closed.
		// To break such iteration channel needs to be closed explicitly.
		// Otherwise range would block forever in the same way as for nil channel
	}

	// Using Select Statement with channels
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sendWithQuit(eve, odd, quit)
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("We can use a 'select' statement to read from channel ")
	selectRecieve(eve, odd, quit)

	// comma ok idiom in select statements
	time.Sleep(time.Millisecond * 5000)
	fmt.Println("We can use a comma,ok idiom when reading from a channel")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("This examples uses this within a select statement to check if channel is open")
	//quit channel was closed, we need to create a new channel in its place
	quit = make(chan int)
	go sendWithQuit(eve, odd, quit)
	commaOkSelect(eve, odd, quit)

}

func commaOkSelect(e, o, q <-chan int) {

	for {
		select {
		case v := <-e:
			fmt.Println("value from e:", v)
		case v := <-o:
			fmt.Println("value from o:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("channel now close, value read :", v, ok)
			} else {
				fmt.Println("channel still open : ", ok)
			}
			return
		}
	}
}

func selectRecieve(e, o, q <-chan int) {
	for {
		//Can use select statement for reading from channels!

		select {
		case v := <-e:
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("value read from even channel:", v)
		case v := <-o:
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("value read from odd channel:", v)
		case <-q:
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("value read from quit channel, quiting!")
			return
		}
	}

}

func sendWithQuit(e, o, q chan<- int) {
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(q)
}

func sender(c chan<- int) {
	//send
	for i := 0; i < 6; i++ {
		fmt.Println("adding to channel")
		c <- i
	}
	close(c) //if not closed we run into a deadlock - try commenting this line!
}

func reciever(c <-chan int) {
	//recieve
	fmt.Println(<-c)
}
