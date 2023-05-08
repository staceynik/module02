package main

import "fmt"

func main() {
	var A *int
	B := 5
	A = &B
	fmt.Println("Value of B via pointer A:", *A)
	*A = 10
	fmt.Println("New value of B:", B)
}
