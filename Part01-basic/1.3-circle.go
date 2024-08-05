package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// BOOL PRACTICE

	var age = 24
	var minor = age < 18

	fmt.Printf("I am %v years old, am I a minor? \n%v", age, minor)

	// TIMER

	var cnt = 10
	for cnt > 0 {
		fmt.Println(cnt)
		time.Sleep(time.Second)
		cnt--
	}
	fmt.Println("liftoff!")

	// HOMEWORK : GUESS THE NUMBER

	var num = rand.Intn(100) + 1
	var res = 64

	for {
		fmt.Printf("%v is not the right answer!\n", num)
		num = rand.Intn(100) + 1
		time.Sleep(time.Second)

		if num == res {
			fmt.Printf("%v is the right answer!\n", num)
			break
		}
	}

}
