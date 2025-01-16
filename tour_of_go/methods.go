package main

import "fmt"
import "math"

// let us define a random structure called demostruct

type demostruct struct{
	X,Y float64
}

//a method is a function with a very special component called receiver
//receiver is present between "func" and method name

func (d demostruct) abs() float64{
	return math.Sqrt(d.X*d.X + d.Y*d.Y)
}

func main(){
	d:=demostruct{3, 4}
	fmt.Println(d.abs())
}


//its the exact same thing as providing this as an argument to the function though as shown below
/*
type demostruct struct{
	X,Y float64
}

//a method is a function with a very special component called receiver
//receiver is present between "func" and method name

func abs(d demostruct) float64{
	return math.Sqrt(d.X*d.X + d.Y*d.Y)
}

func main(){
	d:=demostruct{3, 4}
	fmt.Println(abs(d))
}
*/