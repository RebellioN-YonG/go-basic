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
	fmt.Println()
	// HW chess

	var chess [8][8]rune

	chess[0][5] = 'b'
	chess[0][6] = 'n'
	chess[0][7] = 'r'
	chess[0][3] = 'q'
	chess[0][4] = 'k'
	chess[0][0] = chess[0][7]
	chess[0][1] = chess[0][6]
	chess[0][2] = chess[0][5]

	chess[7][5] = 'B'
	chess[7][6] = 'N'
	chess[7][7] = 'R'
	chess[7][3] = 'Q'
	chess[7][4] = 'K'
	chess[7][0] = chess[7][7]
	chess[7][1] = chess[7][6]
	chess[7][2] = chess[7][5]

	for col := range chess[1] {
		chess[1][col] = 'p'
	}

	for col := range chess[6] {
		chess[6][col] = 'P'
	}

	display(chess)
}

func display(board [8][8]rune) {

	for i := 0; i < 8; i++ {
		for col := range board[i] {

			if i > 1 && i < 6 {
				board[i][col] = ' '
			}
			fmt.Printf("%c", board[i][col])
		}
		fmt.Println("")
	}

}
