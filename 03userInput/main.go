package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome here kinshuk"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the project name")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading ", input)

}
