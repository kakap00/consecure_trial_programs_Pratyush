package main

import(
	"fmt"
	"time"
)




func main(){
	tick := time.Tick(100*time.Millisecond)
	boom := time.After(500*time.Millisecond)
	for{        //will run infinitely until it returns and breaks out
		select{ //listens on all channel cases
		case <- tick:
			fmt.Println("tick.")
		
		case <- boom:
			fmt.Println("BOOM!!!")
			return
		
		default:
			fmt.Println("...")
			time.Sleep(50*time.Millisecond)
		}
	}
}


//IMPORTANT*****************