package main

import "fmt"

var x int = 43

func main() {
	fmt.Println(x)
	foo()

}
func foo() {
	fmt.Println(x)
}
