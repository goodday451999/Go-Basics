package main

import (
	"fmt"
)

func main(){
	var array1[5] int
	fmt.Println(array1)  //[0 0 0 0 0]
	
	array1[2] = 3
	fmt.Println(array1)  //[0 0 3 0 0]
	
	array2 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(array2)  //[1 2 3 4 5]
	
	array3 := [] int {1, 2, 3, 4, 5}
	array3 = append(array3, 6) //add an element at end
	fmt.Println(array3)        //[1 2 3 4 5 6]
	
	i := 2 //add an element at index 2
	array4 := make([]int, len(array3) - i)
	copy(array4[i - 2:], array3[i:])
	fmt.Println(array4)  //[3 4 5 6]
	array3[2] = 9
	array3 = append(array3, 0)
	fmt.Println(array3) //[1 2 9 4 5 6 0]
	copy(array3[i + 1:], array4[i - 2:])
	fmt.Println(array3) //[1 2 9 3 4 5 6]
	
	j := 4 //remove the 3rd index element
	copy(array3[j:], array3[j + 1:])
	fmt.Println(array3[:len(array3) - 1])  //[1 2 9 3 5 6]
	
	array3 = array3[:len(array3) - 1]
	fmt.Println(array3)  //[1 2 9 3 5 6]
	array3 = append(array3, array3[1])
	fmt.Println(array3)  //[1 2 9 3 5 6 2]
	
	// remove elements 
 	array3 = append(array3[:2], array3[4:]...)
	fmt.Println(array3)  //[1 2 5 6 2]
	  
}

