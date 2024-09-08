package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PART PRACTICE: TRAVEL TO MAR
func main() {
	line := rand.Intn(3) + 1
	speed := rand.Intn(15) + 16
	trip := rand.Intn(2) + 1
	price := rand.Intn(50) + 36
	days := 621000 / (speed * 36 * 24)
	var triptype = "One-Way Trip"
	var name = "SpaceX"
	fmt.Printf("%-18v %-5v %-15v %v\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("================================================")

	for cnt := rand.Intn(10) + 1; cnt > 0; cnt-- {

		fmt.Printf("%-18v %-5v %-15v %-4vM$\n", name, days, triptype, price)
		time.Sleep(time.Second)

		speed = rand.Intn(15) + 16
		days = 621000 / (speed * 36 * 24)
		price = rand.Intn(50) + 36
		switch line {
		case 2:
			name = "Virgin Galactic"

		case 3:
			name = "Space Adventure"
		}

		switch trip {
		case 2:
			triptype = "Round-Trip"
			price = price * 2
		}

	}

}
