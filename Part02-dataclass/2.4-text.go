package main

import (
	"fmt"
	// "math/big"
	// "math"
	// "math/rand"
	// "unicode/utf8"
)

func main() {

	fmt.Printf("1\n")
	fmt.Printf(`1\n
	`)
	fmt.Printf(`THE TIME HAS COME!
	I WANT TO SEX!
	WXCB!
	`)

	fmt.Printf(`C:\GO
	`)

	var pi rune = 960

	var c rune = 8964

	fmt.Printf("%c %c\n", pi, c)

	message := "weinnie"

	for i := 0; i < len(message); i++ {
		fmt.Printf("%c\n", message[i])
	}
	fmt.Println("")

	// caesar cipher: ASCII + 1
	c1 := 'a'
	c1 += 1
	fmt.Printf("%c\n", c1)

	// ROT13: ASCII + 13
	msg := "the time has come! i want to sex!"

	for i := 0; i < len(msg); i++ {
		c2 := msg[i]
		if c2 >= 'a' && c2 <= 'z' {
			c2 += 13
			if c2 > 'z' {
				c2 -= 26
			}
		}
		fmt.Printf("%c", c2)
	}
	fmt.Println("")

	// UTF-8
	fmt.Println(len(msg), "bytes")
	fmt.Println(utf8.RuneCountInString(msg), "runes")

	msg2 := "到点了! 我想cb!"

	fmt.Println(len(msg2), "bytes")
	fmt.Println(utf8.RuneCountInString(msg2), "runes")

	for i, c := range msg2 {
		fmt.Printf("%v  %c\n", i, c)
	}

	// HOMEWORK: ENCRYPTION AND DECRYPTION
	s1 := "L fdph, L vdz, L frqtxhuhg."
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			c -= 3
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
	s2 := "Hola Estación Espacial Internacional"

	for _, c := range s2 {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}

		fmt.Printf("%c", c)
	}
}
