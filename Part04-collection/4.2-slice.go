package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sex := [...]string{
		"ddl!", "wxcb!", "cb!", "sex!",
		"the", "time", "has", "come",
	}

	fmt.Println(sex[3:])

	// i in slice[i] is byte, not rune or int
	cb := "到点了！i want2sex!"
	fmt.Println(cb[:6])

	lang := []string{
		"c", "cpp", "java", "python",
		"javascript", "golang", "rust",
		"c#",
	}

	sort.StringSlice(lang).Sort()
	fmt.Println(lang)
	fmt.Println(strings.Join(lang, "|"))
	for i := 0; i < len(lang); i++ {
		fmt.Println("NEW " + lang[i])
	}

}
