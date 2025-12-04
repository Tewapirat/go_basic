package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {


	var wg sync.WaitGroup
	wg.Add(3)

	for i := range 3 {
		go func (i int, wg *sync.WaitGroup)  {
			
			defer wg.Done()
			
			fmt.Println(i)
			
		}(i,&wg)
	}

	wg.Wait()



	
// 	fmt.Println("hello main thread 1")
// 	time.Sleep(2 * time.Second)
// 	go hello()
// 	fmt.Println("Hello main thread 2")
	
// }

// func hello(){
// 	for {
// 		fmt.Println("hello form another thread")
// 		time.Sleep(200 * time.Millisecond)
// 	}
}
