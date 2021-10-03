package main

import "fmt"

// variable declaration outside the function should start with var
// var (
// 	str string = "str"
// 	num int = 123
// 	isTrue bool = true
// )

func main() {
	// fmt.Println("Hello go variables")

	// Individual variable declaration
	// var name string
	// var age int
	// var location string
	
	// variable assignment 
	// name = "Sajan"
	// age = 30
	// location = "india"

	// declare list of variables
	// var (
	// 	name string
	// 	age int
	// 	location string
	// )

	// or

	// var (
	// 	name, location string
	// 	age int
	// )

	// declaration, initialization on the go
	// var (
	// 	name string = "sajan"
	// 	age int = 30
	// 	location string = "India"
	// )

	// or

	// var (
	// 	name = "sajan"
	// 	age = 30
	// 	location = "India"
	// )
	
	// or
	// var (name, age, location = "sajan", 30, "india")
	

	// sorthand variable declaration with implicit type
	name, age, location := "sajan", 30, "india"
	
	fmt.Println(name, age, location)


	greet := func (){
		fmt.Println("hello", name)
	}

	greet()
}

