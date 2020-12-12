package main

import (
	"fmt"
)

func doSum(x int, y int) int {
	return x + y
}

func main(){
	x := 2
	y := 3
	sum := doSum(x, y)
	fmt.Println("Sum: ", sum) 	 
}

