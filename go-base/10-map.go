package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int)
	
	m["abc"] = 1
	m["def"] = 2
	m["ghi"] = 3
	
	fmt.Println(m)
	fmt.Println(m["def"])
	
	delete(m, "def")
	fmt.Println(m)
	  
}


