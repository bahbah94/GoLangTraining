package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wq sync.WaitGroup

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	wq.Add(2)
	go foo()
	go bar()
	wq.Wait()
}


func foo(){
	for i := 0; i < 15 ; i++{
		fmt.Println("Foo :", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wq.Done()
}

func bar(){
	for i:= 0; i < 15 ; i++{
		fmt.Println("Bar : ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wq.Done()
}

