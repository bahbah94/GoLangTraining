package main

import "fmt"

func zero(z *int) {// so here we are receiving a pointer to a int
	*z = 0 // and then i deferenced it and give a value of zero
}

func main() {
	x := 5
	zero(&x) // here we are taking memory of x and passing it to zero function
	fmt.Println(x)
}
