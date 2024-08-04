package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Please write the project name")
	fmt.Println("What that project is about")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	fmt.Println("thanks for letting me know ", input)

	if err != nil {
		fmt.Println("", err)
	} else {
		fmt.Println("all fine")
	}

	// strings

	fmt.Println("check this", strings.TrimSpace(input))

}
