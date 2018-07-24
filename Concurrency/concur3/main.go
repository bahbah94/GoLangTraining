package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wq sync.WaitGroup
var counter int

func main(){
	wq.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wq.Wait()
	fmt.Println("Finla Counter:", counter)
}
func incrementor(s string){
	for i:=0; i< 20; i++{
		x := counter // here is the race condition,
		// both go routines are accessing this global variable and updating it or rather reading it and then doing
		// something
		x++
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
	counter = x
	fmt.Println(s,i,"COunter:", counter)
	}
	wq.Done()
}