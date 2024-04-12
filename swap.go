package main

import "fmt"

func swap() {
	a := 2
	b := 1

	fmt.Println("A = ", a)
	fmt.Println("B = ", b)

	a, b = b, a

	fmt.Println("\nAfter Swap\n")

	fmt.Println("A = ", a)
	fmt.Println("B = ", b)
}
