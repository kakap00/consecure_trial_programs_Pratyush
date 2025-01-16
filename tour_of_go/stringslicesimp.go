package main

import(
	"fmt"
)

func main(){
	s := [4]string{"a","b","c","d"}
	fmt.Println(s)
	slice1 := s[0:2]
	slice2 := s[2:4]
	fmt.Println(slice1,slice2)
	slice1[1] = "x"            //basically, slices are not only representations of the main string
	fmt.Println(s)             //changing a part of a slice changes the parent string as well!!
}
