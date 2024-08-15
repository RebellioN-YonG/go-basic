package main

import (
	"fmt"
)

func main() {

	fmt.Println()

	lang := []string{
		"c", "cpp", "java", "python",
		"golang", "rust",
		"c#",
	}

	lang1 := lang[:6]

	fmt.Println(lang1, len(lang1), cap(lang1))

	lang2 := append(lang1, "ts", "js")

	lang3 := append(lang1, "ts", "js", "php")

	fmt.Println(lang2, lang3, cap(lang2), cap(lang3))

	lang4 := lang[:6:6]
	fmt.Println(lang4, len(lang4), cap(lang4))

	// if we change lang2[1], lang3[1] won't be changed ,they aren't in a same array
	// but if we change lang1[1], lang[1] will change, cuz they are in the same array

	lang2[1] = "r"
	lang1[1] = "r"

	fmt.Println(lang, lang1, lang2, lang3)

	lang4 = append(lang4, "swift")

	fmt.Println(lang, lang4)

	// using make to save ram, 1st para is len, 2nd is cap
	lang = make([]string, 0, 10)

	dump("lang", lang)

	lang = append(lang, "c", "cpp", "js")

	dump("lang", lang)

	//  decalre func that has variable para
	tw := tera("new", "venus", "mars")
	fmt.Println(tw)

	planets := []string{"enus", "mars", "jupiter"}

	np := tera("new", planets...)

	fmt.Println(np)

}

func dump(label string, s []string) {
	fmt.Printf("%v :length %v, capacity %v\n%v\n", label, len(s), cap(s), s)

}

// ... means para is variable
func tera(pre string, worlds ...string) []string {

	new := make([]string, len(worlds))

	for i := range worlds {
		new[i] = pre + " " + worlds[i]
	}

	return new
}
