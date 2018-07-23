package main

import "fmt"

func main(){
	records := make([][]string, 0)
	student1 := make([]string,4)
	student1[0] = "Foster"
	student1[1] = "Randall"
	student1[2] = "100"
	student1[3] = "75"
	// store the record
	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "GOmez"
	student2[1] = "lisa"
	student2[2] = "76"
	student2[3] = "89"

	records = append(records, student2)
	fmt.Println(records)
}
