package main

import (
	"fmt"
	"sync"
	"time"
)

func take_workers(id int, wg *sync.WaitGroup){
	defer wg.Done()   //we will defer this statement until everything finishes running
	fmt.Printf("worker %d has started\n", id)
	time.Sleep(time.Second*5)
	fmt.Printf("worker %d has ended\n", id)
}


func main() {
	var wg sync.WaitGroup

	for i:=1;i<=5;i++{
		wg.Add(1)    //so you increment the counter by 1 everytime for loop runs here
		go take_workers(i, &wg)
	}
	wg.Wait()        //wait will wait for Done() to be called. only then will println print
	fmt.Println("all workers are done!")
}