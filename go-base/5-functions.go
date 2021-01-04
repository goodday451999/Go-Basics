package main

import (
	"fmt"
)

func op(x, y int) (int, int) {
	var output1 = x + y
	var output2 = y - x
	return output1, output2
}
func doSum(x int, y int) int {
	return x + y
}

func main(){
	x := 2
	y := 3
	sum := doSum(x, y)
	fmt.Println("Sum: ", sum) 
	
	add, sub := op(x, y)
	fmt.Println("Add: ", add)
	fmt.Println("Sub: ", sub)
}

