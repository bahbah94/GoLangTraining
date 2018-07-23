package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "good morning", // never forget the commas
		1: "Kumno?",
		2: "Namaste",
		3: "Heya",
	}
	fmt.Println(myGreeting)
	if val, exists := myGreeting[2]; exists{
		delete(myGreeting,2)
		fmt.Println("val: ", val)
		fmt.Println("Exisits: ", exists)
	}else {
		fmt.Println("The value doesnt exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
}