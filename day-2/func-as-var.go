package main

import "fmt"

func main(){
	// Function as variables

	// We can use this to prototype functions before actuallly declaring them

	var v func(int) int
	v = func(x int) int {return x*x}

	result := v(5)
	fmt.Println(result)
}
