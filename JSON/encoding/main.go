package main

import (
	"encoding/json"
	"os"
)

type Person struct{
	First string
	Last string
	Age int
	notExported int
}

func main(){
	p1 := Person{"Rad","Brad", 21, 77}
	json.NewEncoder(os.Stdout).Encode(p1)
}
