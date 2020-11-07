package main

import "fmt"

func main() {
	//define map
	// emails:= make(map[string]string) //first string is data type of key, second is data type of value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"
	// fmt.Println(emails["Bob"])
	// fmt.Println(len(emails))
	// fmt.Println(emails)

	// delete(emails, "Bob")
	// fmt.Println(emails)
	

	// declare map and key value
	emails := map[string] string {"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))
	fmt.Println(emails)
}