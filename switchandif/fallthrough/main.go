package main

import "fmt"

func main(){
	switch "Marcus"{
	case "Tim":
		fmt.Println("Wassup tim")
	case "Jenny":
		fmt.Println("Wassup jenny")
	case "Marcus":
		fmt.Println("Wassup marcus")
		fallthrough // runs the case next after it matches.
	case "Cuck":
		fmt.Println("Wassup cuck")
	case "John":
		fmt.Println("Wassup john")
		fallthrough // does not runthrough because doesnt match with previous case
	case "Julian":
		fmt.Println("Wassup hello")
	}
}
