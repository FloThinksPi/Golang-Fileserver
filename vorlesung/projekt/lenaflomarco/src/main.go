package main

import (
	"fmt"
)

//main - main function
func main() {
	fmt.Print("test")
	testfunction(1)
}

//testfunction - function to test for unit tests
func testfunction(f int) int {
	return f
}
