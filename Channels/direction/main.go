package main

import "fmt"

func main(){
	c := incrementor()
	cSum := puller(c)
	for n:= range cSum{ // we can just put puller on the range
		fmt.Println(n)
	}
}
func incrementor() <-chan int{
	out := make(chan int)
	go func(){
		for i := 0; i < 10; i++{
			out <- i
		}
		close(out)
	}()
	return out
}
func puller(c <-chan int) <-chan int{ // <-cha int we can only receive things from a channel
	out := make(chan int)
	go func(){
		var sum int
		for n := range c{
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

