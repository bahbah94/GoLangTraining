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
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
	counter = x
	fmt.Println(s,i,"COunter:", counter)
	}
	wq.Done()
}