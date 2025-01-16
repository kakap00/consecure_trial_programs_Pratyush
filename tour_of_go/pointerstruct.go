package main

import(
	"fmt"
)


type structpointers struct{
	X int
	Y int
}


func main(){
	v:=structpointers{2,3}
	structpointer := &v
	structpointer.X = 5912874092
	fmt.Print(v)
}