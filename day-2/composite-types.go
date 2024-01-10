package main

import "fmt"

func main(){

	// Arrays

	// fixed sized , contain same type of values

	var prices [3]int = [3]int {1,2,3}
	
	fmt.Println(prices)

	// Slices

	// same as arrays but size can be dynamic

	items := []int {1,2,3}
	fmt.Println(items)

	// Structs

	// Custome data types

	// data type declaration
	type Person struct{
		name string
		age int
	}

	// oobject delaration

	var umair Person = Person{"Umair",100}
	fmt.Println(umair)
	fmt.Println(umair.name)
	fmt.Println(umair.age)

	// Maps

	// maps are key value pairs

	var hights map[string]int = map[string]int{"umair":20,"ali":30}
	fmt.Println(hights)
	fmt.Println(hights["umair"])

}
