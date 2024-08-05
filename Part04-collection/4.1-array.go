package main

import "fmt"

func main() {

	var sex [4]string
	sex[0] = "w"
	sex[1] = "x"
	sex[2] = "c"
	sex[3] = "b"

	fmt.Println(sex)

	a := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	fmt.Printf("%d\n", a)

	for i := 0; i < len(sex); i++ {
		fmt.Println(i, sex[i])
	}

}
