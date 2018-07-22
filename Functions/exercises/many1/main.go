package main

import "fmt"

func main(){
	half := func(n int)(int, bool){ // here we assign func to a variable
		return n / 2, n%2 == 0
	}
	fmt.Println(half(2))
	}

