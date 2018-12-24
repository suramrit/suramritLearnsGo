package main

import "fmt" 

//Recursive Fibonacci in golang .. 

func fib(n int) int{
	if n <1{
		return 0
	} else if n <= 2{
		return 1
	} else{
		v:=fib(n-1) + fib(n-2)
		fmt.Println("The ",n,"-th fibonacci number is",v)
		return v
	}
}

func main(){
		n:=12
		fib(n)
}