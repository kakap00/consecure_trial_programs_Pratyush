package main

import "fmt"


func summerman(int_array []int , channel1 chan int){
	sum := 0
	for _,value := range int_array{
		sum = sum+ value
	}
	channel1 <- sum
}



func main(){
	channel1:= make(chan int)
	int_array := []int{1,2,3,4,5,6}
	go summerman(int_array[len(int_array)/2:],channel1)
	go summerman(int_array[:len(int_array)/2], channel1)
	firstsum := <-channel1                        //CHANNEL IS BLOCKED UNTIL THIS GUY RECEIVES FIRSTSUM!!!
	secondsum := <-channel1

	fmt.Println(firstsum,secondsum)
}

//crazyyyyyyy