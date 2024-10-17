package main

import "fmt" //formatting strings and printing msgs

func main() {
	name := "jake"
	age := 33
	//print
	fmt.Print("Hello, \n")
	fmt.Print("world ")

	fmt.Println("Hello klausese")
	fmt.Println("Goodbuu")

	fmt.Println("My name is", name, "i am", age, "years old")

	//formatted strings
	//PrintF (formatted string) %_ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age)

	fmt.Printf("my name is %q and my age is %q \n", name, age) //quote specifier

	fmt.Printf("my name is %T and my age is %T \n", name, age) //type specifier

	fmt.Printf("you scored %0.1f points \n", 33.2233) //float specifier

	//sprintf save formatted strings
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)

	fmt.Println("the saved string is: \n", str)

}
