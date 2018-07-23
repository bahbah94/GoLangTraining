package main

import "fmt"

type vehicles interface{

}

type vehicle struct{
	Seats int
	MaxSpeed int
	Color string
}
type car struct{
	vehicle
	Wheels int
	Doors int
}
type plane struct{
	vehicle
	Jet bool
}
type boat struct{
	vehicle
	Length int
}

func main(){
	prius := car{}
	tacoma := car{}
	boeing67 := plane{}
	boeing585 := plane{}
	sanger := boat{}
	nautique := boat{}
	rides := []vehicles{prius, tacoma, boeing67, boeing585, sanger, nautique}
	for key, value := range rides{
		fmt.Println(key, " - ", value)
	}
	}
