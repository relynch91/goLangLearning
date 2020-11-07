package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting Method (value receiver)

// p is similiar to this

func (p Person) greet() string { //value receiver
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}


// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string)  {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
	// return
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f",
		age: 25}
	person2 := Person{ firstName: "Bob", lastName: "Johnson", city: "New York", gender: "M",
		age: 30}

	// fmt.Println(person1)
	// fmt.Println(person1.age)
	// person1.age ++
	// fmt.Println(person1.age)

	person2.getMarried("Thompson")
	person2.greet()

	person1.hasBirthday()
	// person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person2.greet())


}
