package main

import (
	"fmt"
)

func inc(x int) {
	x++
}

func incP(x *int) {
	*x++
}

func main(){
	x := 2
	fmt.Println("Input: ", x)
	
	inc(x)
	fmt.Println("Output of inc: ", x)
	
	incP(&x)
	fmt.Println("Output of incP ", x) 	
	
	
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	
}

