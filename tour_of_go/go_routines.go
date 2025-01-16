//go routines are concurrent processes that run alongside the main process/thread
//when we encapsulate a function/method in a go routine, it will execute simultaneously along with the main process
//one thing to note here is the fact that go routines will only run as long as main process runs
//if main process finishes executing, goroutine also stops at whatever point of execution


package main

import (
	"fmt"
	"time")

func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(1000*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("go_routine")
	say("main_thread_routine")
}

//notice sometimes when only 4 instanes of goroutine are printed because main finished first