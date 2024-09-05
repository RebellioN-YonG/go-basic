package main

import (
	"fmt"
	"sort"
)

func main() {

	var nowhere *int
	fmt.Println(nowhere)

	// do not use fmt.Println(*nowhere)
	// which leads to panic

	// to avoid panic, we can use if statement
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	var nobody *person
	fmt.Println(nobody)
	// nobody.birthday() // will cause panic

	var fn func(a, b int) int
	fmt.Println(fn == nil)

	food := []string{"onion", "carrot", "celery", "apple", "soup"}
	sortStrings(food, nil)
	fmt.Println(food)

	// if a slice is declared but not initialized
	// it's value is nil
	var soup []string
	fmt.Println(soup == nil)

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))

	soup = append(soup, "tomato", "onion", "carrot")
	fmt.Println(len(soup), soup)

	// skipping create nil slice,
	// passing nil to a func that returns a non-nil slice
	sopa := mirepoix(nil)
	fmt.Println(sopa)

	// as with slices , a declared map without
	// a composite literal or make() has nil value.

	var soep map[string]int
	fmt.Println(soep == nil)

	// maps can be read when nil,
	// though writing to a nil map will cause panic

	/* measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soep {
		fmt.Println(ingredient, measurement)
	} */
	
	// if a interface type variable is assginged without a value
	// then it's value and type are both nil
	// the variable itself is nil
	var v interface{}
	fmt.Printf("%T %v %v\n",v , v, v == nil)
	
	// both a interface's value and type are  nil,
	// the interface itself equals to nil 
	// if it's value is nil but the type is not
	// then the interface does not equal to nil
	var p *int
	v = p
	fmt.Printf("%T %v %v\n",v , v, v == nil)
	fmt.Printf("%#v\n", v)


 
	n := newNum(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)

}

// if u don't want to use nil,
// just declare a small structure with a few methods

type number struct {
	value int
	valid bool
}

func newNum(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}



func mirepoix(ingredient []string) []string {
	return append(ingredient, "onion", "carrot", "tomato")
}

// if nil is passed for less func,
// if defaults to a func that sorts alphabetically
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, less)
}

type person struct {
	age int
}

func (p *person) birthday() {
	// use if statement to avoid panic
	if p == nil {
		return
	}

	p.age++
}
