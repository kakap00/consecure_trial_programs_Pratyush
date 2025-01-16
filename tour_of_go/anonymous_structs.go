package main

import(
	"fmt"
)

func main() {
	type STRUCT_TYPE struct{
		X int
		Y bool
	}

	anonymous_struct := []struct{      //[] defines the fact that its a slice struct. we define slice inside
		A int
		B int
	}{
		{A:3, B:4},
	}

	s:=STRUCT_TYPE{3,true}

	fmt.Println(s,anonymous_struct)
}