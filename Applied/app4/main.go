package main

import "fmt"

func main(){
	//setting the pipeline
	c := gen(2,3)
	out := sq(c)

	//consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)
}
func gen(num ...int) chan int{
	out := make(chan int)
	go func(){
		for _, n := range num{
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

