package main

import (
	"fmt"
	"sync"
	"time"
)


var (
	mu		sync.Mutex
	balance int = 1000
)

func main(){

	
	doneCh := make(chan bool,3)


	go updateBlance(doneCh, 100)
	go updateBlance(doneCh, 150)
	go updateBlance(doneCh, 200)

	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println(balance)
}

func updateBlance(doneCh chan<- bool , amount int){

	mu.Lock()

	fmt.Println("Updating Blance")
	time.Sleep(1 * time.Second)
	balance -= amount
	doneCh <- true


	mu.Unlock()
	fmt.Println("Blance Updated")
}