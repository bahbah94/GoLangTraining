package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) fullName() string{ // p person is a receiver, so ,
// so any value of the type person can call this method. cool
	return p.first + " " + p.last
}

func main(){
	p1 := person{"Randall", "Dkhar", 41}
	p2 := person{"Rambo", "Killer", 54}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}