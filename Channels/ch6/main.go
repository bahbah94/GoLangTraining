package main

import "fmt"

// here multiple functions reveive from channel
func main(){

	c := make(chan int)
	done := make(chan bool)

	go func(){
		for i := 0; i < 1000 ; i++{
			c <- i
		}
		close(c)
	}()

	go func(){
		for n := range c{
			fmt.Println(n)
		}
		done <- true
	}()
	go func(){
		for n := range c{
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}
