package main

import "fmt"

func max(n ...int) int{
	var largest int
	fmt.Printf("%T \n", n)
	for _, v := range n { // here the '_' stands for the index
		if v > largest {
			largest = v
		}
	}
	return largest
}
 func main(){
 	q := max(4,6,7,8,9)
 	fmt.Println(q)
 }