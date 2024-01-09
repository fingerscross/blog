package main

import "fmt"

func QuickSort(num []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := num[i]
	for i < j {
		for i < j && num[i] < p {
			i++
		}
		for i < j && num[j] > p {
			j--
		}
		num[i], num[j] = num[j], num[i]

	}
	num[i] = p

	QuickSort(num, left, i-1)
	QuickSort(num, i+1, right)
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	n := len(a) - 1
	QuickSort(a, 0, n)
	fmt.Println(a)
}
