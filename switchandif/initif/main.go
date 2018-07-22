package main

import "fmt"

func main(){

	b := true

	if food := "chocolate"; b{// we are initializing here to keep scope tight,
	// and using the variable b to check if condition is true i.e if food variable set to chocolate exists
		fmt.Println(food)
	}
}
