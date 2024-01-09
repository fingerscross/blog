package main

import (
	"fmt"
)

func maopao(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
			}
			a[j], a[i] = a[i], a[j]
		}
	}
	return
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	maopao(a)
	fmt.Print(a)
}
