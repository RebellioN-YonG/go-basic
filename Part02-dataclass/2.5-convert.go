package main

import (
	"fmt"
	"strconv"
	// "math/big"
	// "math"
	// "math/rand"
	// "unicode/utf8"
)

func main() {

	age := 24
	agef := float64(age)
	fmt.Println(agef + 1.0)

	var pi rune = 960
	fmt.Print(string(pi))

	cd := 9

	str := fmt.Sprintf("Launch in T mins %v secs.", cd)
	fmt.Println(str)

	// Itoa: Interger to ASCII
	str1 := "Launch in T mins" + strconv.Itoa(cd) + " secs."
	fmt.Println(str1)

	// Atoi: ASCII to Interger

	cd1, err := strconv.Atoi("10")
	if err != nil {

	}
	fmt.Println(cd1)

	launch := false

	launchtxt := fmt.Sprintf("%v", launch)
	fmt.Println("Ready 4 launch:", launchtxt)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready 4 launch:", yesNo)

	// TEST: BOOL 2 INT

	check := false
	var yon int

	switch check{
		case true:
			yon = 1
		case false:
			yon = 0
	}
	fmt.Println(yon)

	// HOMEWORK: STRING 2 BOOL
	var cu bool
	yOn := "1"
	switch yOn {
		case "1", "yes", "true":
			cu = true
		case "0", "no", "false":
			cu = false
		}
	
	fmt.Println(cu)		


}
