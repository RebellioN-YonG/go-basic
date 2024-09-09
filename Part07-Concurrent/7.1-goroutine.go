package main

import (
	"fmt"
	"time"
)

func main() {

	/* for i := 0; i < 5; i++ {
		go cb(i)
	}
	fmt.Println()
	*/

	// using channels to communicate between goroutines
	c := make(chan int)

	for i := 0; i < 5; i++ {

		go cb1(i, c)
	}
	/*
		for i := 0; i < 5; i++ {
			cbid := <-c
			fmt.Println("cbid: ", cbid, " has finished!")
		}
	*/
	// using select and time.After to wait for the channel
	// from the cb and the timeout channel
	timeout := time.After(3 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case cbid := <-c:
			fmt.Println("cbid: ", cbid, " has finished!")
		case <-timeout:
			fmt.Println("timeout!")
			return
		}
	}

}

/*
	func cb(id int) {
		time.Sleep(3 * time.Second)
		fmt.Println("...", id, " cb!...")
	}
*/
func cb1(id int, c chan int) {
	// time.Sleep(6 * time.Second)
	time.Sleep(2 * time.Second + 999 * time.Millisecond)
	fmt.Println("...", id, " cb!...")
	c <- id
}
