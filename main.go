package main

import "fmt" //formatting strings and printing msgs

func main() {
	//fmt.Println("Hello world")

	//strings
	var nameOne string = "jude"
	var nameTwo = "Mike"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Kemi"
	nameThree = "kome"

	//fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Friday"
	fmt.Println(nameFour)

	//int
	var age int = 30

	var ageTwo = 21

	ageThree := 33

	fmt.Println(age, ageTwo, ageThree)

	var numOne int8 = 57
	var numthree float64 = 1.2343334

	fmt.Println(numOne, numthree)

}
