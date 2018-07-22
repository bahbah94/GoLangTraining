package main

import "fmt"

func main(){
	n := average(54,67,87,32)
	fmt.Println(n)
}
func average(sf ...float64) float64{
	fmt.Println(sf)
	fmt.Printf("%T", sf)
	total := 0.0
	for _, v := range sf{
		total += v
	}
	return total / float64(len(sf))
}
