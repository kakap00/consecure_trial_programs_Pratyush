package main

import(
	"fmt"
)


type ABC struct{
	X int
	Y int
}


func main(){

	a := ABC{1,2}
	a.X = 12423
	fmt.Println(a.Y)
	fmt.Println(a.X)
	fmt.Printf("this is X %d this is Y %d", a.X, a.Y)
}