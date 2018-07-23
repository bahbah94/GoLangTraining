package main

import (
	"math"
	"fmt"
)

type circle struct{
	radius float64
}

type shape interface{
	area() float64
}

func (c *circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func info(s shape){
	fmt.Println("area is ", s.area())

}

func main(){
	c := circle{5}
	info(&c)// this works if c is a pointer value, e.g &c
}
