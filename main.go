package main

import "fmt"

func main() {
	//var ages = [3]int{22, 23, 21}

	name := [4]string{"jake", "make", "hhs", "yimde"}

	fmt.Println(name)

	//slices
	var scores = []int{100, 45, 67, 54}
	scores[2] = 44
	scores = append(scores, 77)
	fmt.Println(scores)

	//slice ranges
	rangeOne := name[1:3]

	fmt.Println(rangeOne)
}
