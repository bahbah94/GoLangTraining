package main

import "fmt"

func half(n int)(int, bool){
	return n / 2, n%2 == 0
}

func main(){
	h, even  := half(5) // assignns the two returns to the two variables
	fmt.Println(h, even)
}
