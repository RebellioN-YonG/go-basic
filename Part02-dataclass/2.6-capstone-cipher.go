package main

import (
	"fmt"
	"strings"
	// "math/big"
	// "math"
	// "math/rand"
	// "unicode/utf8"
)

// Part02 Capstone2 The Vigenère cipher

func main() {

	// 1 decrypt text with Vigenère
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	originText := ""
	j := 0
	for i := 0; i < len(cipherText); i++ {

		c := cipherText[i] - 'A'
		k := keyword[j] - 'A'
		c = (c-k+26)%26 + 'A'
		originText += string(c)
		j++
		j %= len(keyword)

	}

	fmt.Println(originText)

	// 2 encrypt text with Vigenère
	plainText := "your message goes here"
	j = 0
	pt := strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	ct := ""
	for i := 0; i < len(pt); i++ {
		if pt[i] >= 'A' && pt[i] <= 'Z' {
			c := pt[i] - 'A'
			k := keyword[j] - 'A'
			ct += string((c+k+26)%26 + 'A')
			j++
			j %= len(keyword)
		}
	}
	fmt.Println(ct)

}
