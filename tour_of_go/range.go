package main

import "fmt"

func main(){
	s := []int{1,2,3,4,5,6,7,8}
	for index,element := range s{   //it is important to note here, that range is an OPERATOR 
		fmt.Println(index,element)  //on an array S. range s will simply provide index value pairs
	}                               //until the final element.
}                                   //index,element:= simply assigns index and element to range output
                                    //at every iteration