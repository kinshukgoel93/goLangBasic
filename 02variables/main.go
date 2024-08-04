package main

import "fmt"

// if you want to make public . Make first letter captial L

func main() {

	var username string = "kinshuk"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLogged bool = false
	fmt.Printf("Variable is of type : %T \n", isLogged)

	var smallValue int = 256
	fmt.Printf("Variable is of type : %T \n", smallValue)

	//Default values or aliases

	var anotherVariable int
	fmt.Println("an", anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// no var style
	//    :=
	// allowed only in function

}
