package main

import "fmt"

func main() {
	// Arrays have to be fixed length and name the types
	//slice is an array without a fixed type
	// var fruitArray [2]string
	
	//assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"
	
	// Declare and Assign
	// fruitArray := [2]string{"Apple", "Orange"}
	fruitSlice := []string{"Apple", "Orange", "Banana"}
	
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(len(fruitSlice[1:2]))
	fmt.Println((fruitSlice[1:3]))
}
