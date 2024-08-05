package main

import "fmt"

type cel float64
type fah float64

func c2f(c cel) fah {

	return fah((c * 9.0 / 5.0) + 32)

}

func f2c(f fah) cel {
	return cel((f - 32) / 9.0 * 5.0)

}

func main() {

	line := "====================="
	table := "| %-8v| %-8v|\n"

	var c cel = -40.0
	var f fah = -40.0

	fmt.Println(line)
	fmt.Printf(table, "C", "F")
	for i := 0; i <= (100+40)/5; i++ {

		fmt.Printf(table, fmt.Sprintf("%.1f", c), fmt.Sprintf("%.1f", c2f(c)))
		c += 5.0
	}
	fmt.Println(line)

	fmt.Println(line)
	fmt.Printf(table, "F", "C")
	for i := 0; i <= (100+40)/5; i++ {

		fmt.Printf(table, fmt.Sprintf("%.1f", f), fmt.Sprintf("%.1f", f2c(f)))
		f += 5.0
	}
	fmt.Println(line)

}
