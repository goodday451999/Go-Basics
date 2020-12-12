package main

import (
	"fmt"
)

func main(){
	var array1[5] int
	fmt.Println(array1)
	
	array1[2] = 3
	fmt.Println(array1)
	
	array2 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(array2)
	
	array3 := [] int {1, 2, 3, 4, 5}
	array3 = append(array3, 6)
	fmt.Println(array3)  
}

