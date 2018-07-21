package main

import "fmt"

const meterToYards float64 = 1.095858

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * meterToYards
	fmt.Println(meters, "meters is ", yards, "yards.")
}
