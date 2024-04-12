package main

import "fmt"

func sort() {
	// a := []int{16, 2, 33, 9}
	a := []string{"Hello", "Apple", "Cat", "Frog"}
	n := len(a)
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
