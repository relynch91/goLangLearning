package main

import "fmt"

func main() {

	ids := []int { 31, 34, 45, 64, 66, 45 }

	//with index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0 

	for _, id := range ids {
		sum += id
	}

	// fmt.Println("sum", sum)

	// Range with map
	emails := map[string] string {"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name: " + k)
	}
	
}
