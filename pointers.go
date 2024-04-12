package main

import "fmt"

func pointers() {
	a := 45
	updateVariable(&a)
	fmt.Print("New Value = ", a)
}

func updateVariable(a *int) {
	fmt.Println("Old Value =", *a)
	*a = 10
}
