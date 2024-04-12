package main

import "fmt"

func add() {
	var n1 int
	var n2 int
	var sum int

	fmt.Println("Enter 1st num :")
	fmt.Scanln(&n1)

	fmt.Println("Enter 2nd num :")
	fmt.Scanln(&n2)

	sum = n1 + n2
	fmt.Println("Sum = ", sum)
}
