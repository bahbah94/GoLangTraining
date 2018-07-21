package main

import "fmt"

func main() {
	a := 54

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}
