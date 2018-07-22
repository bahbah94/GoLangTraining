package main

import "fmt"

func main(){
	i := 0
	for{
		fmt.Println(i)
		if i >= 10{ // so this is a condition, ignores the break, so it will print 10
			break
		}
		i++
	}
}
