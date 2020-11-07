package main
import "fmt"

var age int = 28

func main() {
	// create variable using var, const works the same as javascript
	// var name string = "Ryan"
	//shorthand

	name:= "Ryan"
	
	var isCool = true
	isCool = false


	fmt.Println(name, age, isCool);
	fmt.Printf("%T\n", isCool)
}