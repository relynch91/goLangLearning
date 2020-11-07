package main

import "fmt"

func main() {
	//pointer allows you to point to memory address or allocation for the variable
	// pointers allow a program to pass the adress to the data in a function,
	//instead of passing the entire object itself. Saves space BRUH

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // -> *int * designates a pointer
	fmt.Println(*b)
}
