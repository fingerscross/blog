package main

import "fmt"

func separate(a []int) {
	n := len(a)
	slow := 0
	fast := 1
	for fast < n {
		if a[slow] == 0 && a[fast] != 0 {
			a[slow], a[fast] = a[fast], a[slow]
			slow++
			fast++
		} else if a[slow] != 0 && a[fast] != 0 && slow != fast {
			slow++
		} else if a[slow] == 0 && a[fast] == 0 && slow != fast {
			fast++
		} else if slow == fast {
			fast++
		}

	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 0, 6, 7, 9}
	separate(a)
	fmt.Println(a)
}
