package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []string{"zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	//
	//sort.Sort(sort.StringSlice(s))
	//fmt.Println(s)
	//t := sort.StringSlice(s)
	//fmt.Printf("%T \n", t)
	//fmt.Printf("Sorted :%T \n", sort.StringSlice(s))
	//fmt.Printf("Reverse %T \n", sort.Reverse(sort.StringSlice(s)))
	//i := sort.Reverse(sort.StringSlice(s))
	//fmt.Println(i)
	}
