package main

import "fmt"

func main(){
	// Variable types in go
	
	// Basic types
	
	// Integer types 

	// Signed Int

	var a int8 = 1
	var b int16 = 10
	var c int64 = 100

	fmt.Println(a,b,c)
	
	// Unsigned Integers

	var aa uint8 = 255
	var bb uint64 = 1221

	fmt.Println(aa,bb)

	// Machine dependent types

	var aaa int = 100
	var bbb uint = 1000
	var ccc uintptr 

	fmt.Println(aaa,bbb,ccc)


	// Float types

	var i float32 = 3.111
	var j float64 = 2323.23232

	fmt.Println(i,j)

	// Complex numbers 

	// comples64 -> float32 for real and imaginary part
	// comples128 -> float64 for real and imaginary part

	var com complex128 = 1 + 2i
	
	fmt.Println(com)

	// Booleans

	var isTrue bool = true
	var isFalse bool = false 

	fmt.Println(isTrue,isFalse)

	// Strings

	name := "Umair"

	fmt.Println(name)

}
