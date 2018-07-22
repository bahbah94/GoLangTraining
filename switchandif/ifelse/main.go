package main

import "fmt"

func main(){
	if false{
		fmt.Println("firts print statement")
	}else if true{
		fmt.Println("second print statement")
	} else {
		fmt.Printf("thirds print statement")
	}
}
