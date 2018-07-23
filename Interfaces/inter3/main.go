package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()// here we convert s to stringslice and we implement sort
	// sort.Sort(sort.StringSlice(s)) while here stringslice implements interface interface,
	// and we so take an interface as an argument, hence we can use the Sort
	fmt.Println(s)
}
