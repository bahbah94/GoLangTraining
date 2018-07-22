package main // a rune is an alias for int32
import "fmt"

// an integer value identifying a unicode code point
func main(){
	for i := 50; i <= 140; i++{
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
