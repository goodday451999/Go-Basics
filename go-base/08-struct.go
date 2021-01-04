package main

import (
	"fmt"
)

type student struct {
	name string
	age int
} 

func main(){
	s := student{"abc xyz", 10}
	fmt.Println(s)
	fmt.Println(s.age)
}

