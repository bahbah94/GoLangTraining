package main

import "fmt"

type person struct{
	first string
	second string
	age int
}

type DoubleZero struct{
	person
	First string
	LicenseToKill bool
}
func main() {
	p1 := DoubleZero{
		person: person{
			first:  "Randall",
			second: "Lol",
			age:    54,
		},
		First:         "zo yolo",
		LicenseToKill: true,
	}
	p2 := DoubleZero{
		person: person{
			first:  "Loser",
			second: "Loler",
			age:    545,
		},
		First:         "zo mogfo yolo",
		LicenseToKill: false,
	}
	fmt.Println(p1.First, p1.person.first)
	fmt.Println(p2.LicenseToKill, p2.person.age)
}