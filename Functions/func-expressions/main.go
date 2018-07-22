package main

import "fmt"

func main(){
	greeting := func(){ // this is an anonymous function, with no identifier,
	// and assigning it to a variable, and this is known as function expression
		fmt.Println("hello motherfuckers")
	}
}
