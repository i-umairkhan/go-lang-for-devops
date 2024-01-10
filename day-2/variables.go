package main

import "fmt"

func main() {
	// Variable declaration in go

	// 1- Using var key word
	var name string
	var age int

	// 2- Declaring multiple variables
	var (
		firstName, lastName string
		age1                int
	)

	// With values
	var name2 string = "john"
	var age2 int = 10

	// Type Interface
	var name3 = "john"
	var age3 = 30

	// Short variable declaration (Can be used only inside the functions)
	name4 := "john"
	age4 := 10

	fmt.Println(name, age, firstName, lastName, age1, name2, age2, name3, age3, name4, age4)
}
