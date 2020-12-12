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
}

