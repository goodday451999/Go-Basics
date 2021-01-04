package main

import (
	"fmt"
)

func main() {
	a := [6]int{8, 4, 5, 2, 4, 5}
	for i, element := range a {
		fmt.Printf("%d : %d\n", i, element)
	}
	for _, element := range a {
		fmt.Println(element)
	}
	// print the duplicate element
	// Method 1
	for i, ele_i := range a {
		for j, ele_j := range a {
			if i < j && ele_i == ele_j {
				fmt.Printf("%d, %d\n", i, ele_i)
			}
		}
	}
	// Method 2
	var b []int = []int{8, 4, 5, 2, 4, 9, 5}
	for i, element := range b {
		for j := i + 1; j < len(b); j++ {
			if element == b[j] {
				fmt.Println(element)
			}
		}
	}
	
} 
	
