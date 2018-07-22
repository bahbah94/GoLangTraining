package main

import "fmt"

func main(){
	m := make([]string, 1, 24)
	fmt.Println(m)// []
	changeMe(m)
	fmt.Println(m) //[Randall]
}

func changeMe(z []string){
	z[0] = "Randall"
	fmt.Println(z)
}
