package main

import "fmt"

type person struct{
	first string
	second string
	age int
}

func main(){
	p1 := person{"Randall", "Kharkrang", 23}
	p2 := person{"Ben", "Shapiro", 33}
	fmt.Println(p1)
	fmt.Println(p2)
}
