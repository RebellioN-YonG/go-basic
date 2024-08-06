package main

import (
	"fmt"
	"math/rand"
)

// HOMEWORK SHOW THE DATE RANDOMLY

var era = "AD"

func main() {
	year := rand.Intn(30) + 2000
	month := rand.Intn(12) + 1
	days := 31

	switch month {
	case 2:
		days = 28
		if year%4 == 0 && year%100 != 0 {
			days = 29
		}

	case 4, 6, 9, 11:
		days = 30
	}

	date := rand.Intn(days) + 1
	fmt.Println(era, year, month, date)
}
