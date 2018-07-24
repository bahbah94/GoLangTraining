package main

import "fmt"

func factorial(n int) <-chan int{
	out := make(chan int)
	go func(){
		if n == 1{
			out <- 1
		}else{
			out <- n * <-factorial(n-1)
		}
	}()
	return out
	}

func main(){
	n := 4
	fmt.Println(<-factorial(n))
}
