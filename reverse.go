package main

import "fmt"

func reverse() {
	var num int
	rev := 0

	fmt.Println("Enter num to reverse: ")
	fmt.Scan(&num)

	for num > 0 {
		r := num % 10
		rev = rev*10 + r
		num = num / 10
	}

	fmt.Println("Reversed num = ", rev)
}
