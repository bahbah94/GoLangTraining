package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wq sync.WaitGroup
var counter int
var mutex sync.Mutex

func main(){
	wq.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wq.Wait()
	fmt.Println("Finla Counter:", counter)
}
func incrementor(s string){
	for i:=0; i< 20; i++{
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s,i,"COunter:", counter)
		mutex.Unlock()
	}
	wq.Done()
}
