package main

import "fmt"

func main() {

	// HW slice append
	slice := make([]int, 0)

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		cp := cap(slice)
		l := len(slice)

		if l%cp == 0 {
			fmt.Println(slice)
			fmt.Println("current capacity is:", cp)
		}
	}
}
