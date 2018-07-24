package main

import "fmt"
// this is a semaphores, so here we stopped using weight groups, and no longer need package sync
func main(){

	c := make(chan int)
	done := make(chan bool)

	go func(){
		for i := 0; i < 10; i++{
			c <- i
		}
		done <- true
	}()

	go func(){
		for i := 0; i < 10; i++{
			c <- i
		}
		done <- true
	}()
// suppose we remove the below function,
// we are blocking the range for loop which reveives values from the channle so we are effectively running the code
// foreever
	go func(){
		<-done
		<-done
		close(c)
	}()
	for n := range c{
		fmt.Println(n)
	}
}
