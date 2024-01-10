package main
import "fmt"

type Person struct {
	name string
	weight int
}

func Increaseweight (person *Person,new_weight int){
 	person.weight = new_weight
}

func main(){

	// Pointers

	// Pointer holds memory address of a variable

	var val int = 10
	var add_val *int = &val
	
	fmt.Println(val)
	fmt.Println(add_val)
	fmt.Println(*add_val)

	// Example 

	var umair Person = Person{"umair",65}

	fmt.Println("Old weight",umair)

	Increaseweight(&umair,70)

	fmt.Println("Old weight",umair)
}
