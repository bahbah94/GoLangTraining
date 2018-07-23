package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "GoodMorning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}
	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}