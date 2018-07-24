package main

import (
	"fmt"
	"sync"
)

func main(){
	in := gen(2,3)

	//Fan out, multiple functions receiving from just one channel
	c1 := sq(in)
	c2 := sq(in)

	//Fan in
	//multiple
	for n := range merge(c1,c2){
		fmt.Println(n)
	}
}

func gen(num ...int) chan int{
	fmt.Printf("Type of nums %T \n", num)
	out := make(chan int)
	go func(){
		for _, n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int{
	out := make(chan int)
	go func(){
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int{
	out := make(chan int)
	var wq sync.WaitGroup
	fmt.Printf("Type of cs : %T\n", cs)
	wq.Add(len(cs))
	for _, c := range cs{
		go func(ch chan int){
			for n := range ch{
				out <- n
			}
			wq.Done()
		}(c)
	}
	go func(){
		wq.Wait()
		close(out)
	}()
	return out
}
