package main

import (
	"fmt"
)

func main(){
	x := 9
	if x % 2 == 0 && x % 3 == 0 {
		fmt.Println("div by 2 and 3")
	} else if x % 2 == 0 {
		fmt.Println("div by 2")
	} else if x % 3 == 0 {
		fmt.Println("div by 3")
	} else {
		fmt.Println("not div by 2 and 3")
	}
	 
}
