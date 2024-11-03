package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends"
	// fmt.Println(strings.Index(greeting, "th"))
	//fmt.Println(strings.Split(greeting, "fr"))

	// fmt.Println(strings.Contains(greeting, "helllo"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //returns a new string doest alter the old one
	// fmt.Println(strings.ToUpper(greeting))

	age := []int{32, 45, 67, 32, 34, 56}
	sort.Ints(age)
	fmt.Println(age)

	index := sort.SearchInts(age, 67)
	fmt.Println(index)

	names := []string{"jake", "jakee", "make"} //returns alphabetical order

	sort.Strings(names)
	fmt.Println(names)

}
