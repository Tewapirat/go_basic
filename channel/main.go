package main

import (
	"fmt"
)

func main (){


	// Test Channal
	// ch := make(chan int)

	// go func(){
	// 	fmt.Println("test 1")
	// 	ch <- 20
	// }()

	// fmt.Println("test 2")
	// v := <- ch

	// fmt.Println(v)



	// Iinfinity loop fix deadlock
	// ch := make(chan int)

	// go func ()  {
	// 	for{
	// 		v := <- ch
	// 		fmt.Println(v)
	// 	}
	// }()

	// ch <- 10
	// ch <- 20
	// ch <- 30
	// ch <- 40


	// Range -> fix deadlock
	// ch := make(chan int)

	// go func ()  {
	// 	for v:= range ch {
	// 		fmt.Println(v)
	// 	}	
	// }()

	// 	ch <- 10
	// 	ch <- 20
	// 	ch <- 30
	// 	ch <- 40
	
	// ch := make(chan int)

	// go func ()  {
	// 	for{
	// 		ch <- 10
	// 		ch <- 20
	// 		ch <- 30
	// 		ch <- 40
	// 	}
	// }()

	// for v:= range ch{
	// 	fmt.Println(v)
	// }



	// Select with Ch
	// ch1 := make(chan int)
	// ch2 := make(chan int)


	// go func ()  {
	// 	ch1 <- 10
	// 	close(ch1)
	
	// }()

	// go func ()  {
	// 	ch2 <- 20
	// 	close(ch2)
	// }()

	// for {
	// 	select {
	// 	case v := <- ch1:
	// 		fmt.Println("Channel 1",v)
	// 	case v := <- ch2:
	// 		fmt.Println("Channel 2",v)
	// 		}
	// }

		// stop infinity loop
	ch1 := make(chan int)
	ch2 := make(chan int)


	go func ()  {
		ch1 <- 10
		close(ch1)
	
	}()

	go func ()  {
		ch2 <- 20
		close(ch2)
	}()

	cch1,cch2 := false,false

	for {
		if cch1 && cch2 {
			break
		}
		select {
		case v, ok:= <- ch1:
			if!ok{
				cch1 = true
				continue
			}
			fmt.Println("Channel 1",v)
		case v , ok:= <- ch2:
			if !ok {
				cch2 = true
				continue
				
			}
			fmt.Println("Channel 2",v)
			}
	}
}