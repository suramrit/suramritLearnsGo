package main 

import("fmt"
)

func main() {
	fmt.Println("hey hargobind, hope you are having a nice day!")
	fmt.Printf("Hi hargobind, I am told you are %d years old\n", 19)
	c := make(chan int,3) //buffered channel 
	//Channels Block! 
	//c <- 32 //This will block.. 
			//until send and recieve at the same time... 
			//or  buffered channel 
	go func() {
		c<-332 // Write to channel .. routine is blocked until its read
	}()
	fmt.Println(<-c) //reads from the channel... once read, goroutine can unblock
	//fmt.Println(<-c)

	//Directional Channels - only send to or recieve from channel

	fmt.Printf("%T\t\n", c)

	only_send_channel:= make(chan<- int) //recv only
	recv_only := make(<-chan int) // send only
	fmt.Printf("%T\t\n", only_send_channel)
	fmt.Printf("%T\t\n", recv_only)
	//recv_only <- 32 //error! 
	go func () {
		only_send_channel <- 223
	}()
	//fmt.Println(<- only_send_channel) //error!

	//Using send and recieve channels 

	//send
	go sender(c)

	//recieve
	c<-39
	//reciever(c)

	//Ranging over a channel -- blocks until the channel is closed

	for v:= range c {
		fmt.Println(v)
	}

	//Select Statement 
	eve := make(chan int)
	odd:= make (chan int)
	quit:= make(chan int)

	//go send(eve, odd, quit)

	//select_recieve(eve, odd, quit) 

	// comma ok idiom in select statements 
	//quit = make(chan bool)
	go send(eve,odd,quit)
	comma_ok_select(eve,odd,quit)


}

func comma_ok_select(e,o,q <-chan int){

	for {
		select{
		case v:= <-e:
			fmt.Println("value from e:",v)
		case v:= <-o:
			fmt.Println("value from o:",v)
		case v,ok := <- q:
			if(!ok){
				fmt.Println("channel now close, value:",v,ok)
			} else {
				fmt.Println( "channel still open : ",v,ok)
			}
			return
		}
	}
}

func select_recieve(e,o,q <-chan int){
	for {
		select {
		case v:= <-e :
			fmt.Println("value from even channel:",v)
		case v:= <-o :
			fmt.Println("value from odd channel:",v)
		case v:= <-q :
			fmt.Println("value from quit channel:",v)
			return
		}
	}

}

func send(e,o,q chan<- int){
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q<-0
	close(q)
}

func sender(c chan<- int){
	//send 
	for i := 0; i < 6; i++ {
		c<-i
	}
	close(c)
}

func reciever(c <-chan int){
	//recieve
	fmt.Println(<-c)
}
