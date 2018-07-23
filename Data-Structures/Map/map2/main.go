package main

import "fmt"

func main(){
	var myMap = map[string]string{
		"Randall" : "cuck",
		"Soldier" : "No",
	}
	myMap["loner"] = "notmr"
	delete(myMap, "Soldier")
	fmt.Println(myMap)
}
