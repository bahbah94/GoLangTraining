package main

import "fmt"

func main(){
	i := 0
	for {
		i++
		if i%2 == 0{
			continue //if true, it will stop and go to top of loop
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
